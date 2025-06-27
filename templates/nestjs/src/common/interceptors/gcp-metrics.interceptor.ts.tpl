import {
  Injectable,
  NestInterceptor,
  ExecutionContext,
  CallHandler,
  Logger,
} from '@nestjs/common';
import { Observable } from 'rxjs';
import { tap, catchError } from 'rxjs/operators';
import { Monitoring } from '@google-cloud/monitoring';
import { Logging } from '@google-cloud/logging';

@Injectable()
export class GCPMetricsInterceptor implements NestInterceptor {
  private readonly logger = new Logger(GCPMetricsInterceptor.name);
  private monitoring: Monitoring;
  private logging: Logging;
  private projectId: string;

  constructor() {
    this.projectId = process.env.GCP_PROJECT_ID || '{{.GCPProject}}';
    if (this.projectId && this.projectId !== '') {
      this.monitoring = new Monitoring();
      this.logging = new Logging({ projectId: this.projectId });
    }
  }

  intercept(context: ExecutionContext, next: CallHandler): Observable<any> {
    const request = context.switchToHttp().getRequest();
    const startTime = Date.now();
    const method = request.method;
    const url = request.url;

    return next.handle().pipe(
      tap((data) => {
        const duration = Date.now() - startTime;
        this.recordMetrics(method, url, 200, duration);
        this.logRequest(method, url, 200, duration);
      }),
      catchError((error) => {
        const duration = Date.now() - startTime;
        const statusCode = error.status || 500;
        this.recordMetrics(method, url, statusCode, duration);
        this.logRequest(method, url, statusCode, duration, error.message);
        throw error;
      }),
    );
  }

  private async recordMetrics(method: string, url: string, statusCode: number, duration: number) {
    if (!this.monitoring || !this.projectId) return;

    try {
      const projectName = this.monitoring.projectPath(this.projectId);
      
      const dataPoint = {
        interval: {
          endTime: {
            seconds: Math.floor(Date.now() / 1000),
          },
        },
        value: {
          doubleValue: duration,
        },
      };

      const timeSeriesData = {
        metric: {
          type: 'custom.googleapis.com/{{.ProjectName}}/request_duration',
          labels: {
            method: method,
            endpoint: url,
            status_code: statusCode.toString(),
          },
        },
        resource: {
          type: 'generic_node',
          labels: {
            location: 'global',
            namespace: '{{.ProjectName}}',
            node_id: process.env.NODE_NAME || 'default',
          },
        },
        points: [dataPoint],
      };

      await this.monitoring.createTimeSeries({
        name: projectName,
        timeSeries: [timeSeriesData],
      });
    } catch (error) {
      this.logger.error('Failed to record metrics:', error);
    }
  }

  private async logRequest(method: string, url: string, statusCode: number, duration: number, error?: string) {
    if (!this.logging || !this.projectId) {
      console.log(JSON.stringify({ method, url, statusCode, duration, error }));
      return;
    }

    try {
      const log = this.logging.log('{{.ProjectName}}-requests');
      const metadata = {
        resource: {
          type: 'generic_node',
          labels: {
            location: 'global',
            namespace: '{{.ProjectName}}',
            node_id: process.env.NODE_NAME || 'default',
          },
        },
        severity: statusCode >= 400 ? 'ERROR' : 'INFO',
      };

      const entry = log.entry(metadata, {
        method,
        url,
        statusCode,
        duration,
        error,
        timestamp: new Date().toISOString(),
      });

      await log.write(entry);
    } catch (logError) {
      this.logger.error('Failed to write log:', logError);
    }
  }
}
