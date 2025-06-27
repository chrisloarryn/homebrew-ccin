package cmd

import (
	"fmt"
	"path/filepath"
)

const (
	flagGCPProject = "gcp-project"
	domainFormat   = "Domain: %s\n"
	gcpFormat      = "GCP Project: %s\n"
)

func generateNestJSProject(projectName, domainName, gcpProject string) error {
	if err := createProjectDir(projectName); err != nil {
		return err
	}

	// Set default domain if not provided
	if domainName == "" {
		domainName = "item"
	}

	// Generate package.json
	packageJSON := replaceTemplateVars(`{
  "name": "{{PROJECT_NAME}}",
  "version": "0.0.1",
  "description": "{{PROJECT_NAME}} - NestJS CRUD API",
  "author": "",
  "private": true,
  "license": "UNLICENSED",
  "scripts": {
    "build": "nest build",
    "format": "prettier --write \"src/**/*.ts\" \"test/**/*.ts\"",
    "start": "nest start",
    "start:dev": "nest start --watch",
    "start:debug": "nest start --debug --watch",
    "start:prod": "node dist/main",
    "lint": "eslint \"{src,apps,libs,test}/**/*.ts\" --fix",
    "test": "jest",
    "test:watch": "jest --watch",
    "test:cov": "jest --coverage",
    "test:debug": "node --inspect-brk -r tsconfig-paths/register -r ts-node/register node_modules/.bin/jest --runInBand",
    "test:e2e": "jest --config ./test/jest-e2e.json"
  },
  "dependencies": {
    "@nestjs/common": "^10.0.0",
    "@nestjs/core": "^10.0.0",
    "@nestjs/platform-express": "^10.0.0",
    "@nestjs/mongoose": "^10.0.1",
    "@nestjs/swagger": "^7.1.8",
    "@google-cloud/monitoring": "^4.0.0",
    "@google-cloud/logging": "^11.0.0",
    "mongoose": "^7.4.3",
    "class-validator": "^0.14.0",
    "class-transformer": "^0.5.1",
    "reflect-metadata": "^0.1.13",
    "rxjs": "^7.8.1"
  },
  "devDependencies": {
    "@nestjs/cli": "^10.0.0",
    "@nestjs/schematics": "^10.0.0",
    "@nestjs/testing": "^10.0.0",
    "@types/express": "^4.17.17",
    "@types/jest": "^29.5.2",
    "@types/node": "^20.3.1",
    "@types/supertest": "^2.0.12",
    "@typescript-eslint/eslint-plugin": "^6.0.0",
    "@typescript-eslint/parser": "^6.0.0",
    "eslint": "^8.42.0",
    "eslint-config-prettier": "^9.0.0",
    "eslint-plugin-prettier": "^5.0.0",
    "jest": "^29.5.0",
    "prettier": "^3.0.0",
    "source-map-support": "^0.5.21",
    "supertest": "^6.3.3",
    "ts-jest": "^29.1.0",
    "ts-loader": "^9.4.3",
    "ts-node": "^10.9.1",
    "tsconfig-paths": "^4.2.1",
    "typescript": "^5.1.3"
  },
  "jest": {
    "moduleFileExtensions": [
      "js",
      "json",
      "ts"
    ],
    "rootDir": "src",
    "testRegex": ".*\\.spec\\.ts$",
    "transform": {
      "^.+\\.(t|j)s$": "ts-jest"
    },
    "collectCoverageFrom": [
      "**/*.(t|j)s"
    ],
    "coverageDirectory": "../coverage",
    "testEnvironment": "node"
  }
}`, projectName, domainName, gcpProject)

	if err := createFile(filepath.Join(projectName, "package.json"), packageJSON); err != nil {
		return err
	}

	// Generate nest-cli.json
	nestCLI := `{
  "$schema": "https://json.schemastore.org/nest-cli",
  "collection": "@nestjs/schematics",
  "sourceRoot": "src",
  "compilerOptions": {
    "deleteOutDir": true
  }
}`

	if err := createFile(filepath.Join(projectName, "nest-cli.json"), nestCLI); err != nil {
		return err
	}

	// Generate main.ts
	mainTS := replaceTemplateVars(`import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { ValidationPipe } from '@nestjs/common';
import { SwaggerModule, DocumentBuilder } from '@nestjs/swagger';
import { GCPMetricsInterceptor } from './common/interceptors/gcp-metrics.interceptor';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  
  // Global validation pipe
  app.useGlobalPipes(new ValidationPipe({
    whitelist: true,
    transform: true,
    forbidNonWhitelisted: true,
  }));

  // GCP Metrics interceptor
  app.useGlobalInterceptors(new GCPMetricsInterceptor());

  // Swagger configuration
  const config = new DocumentBuilder()
    .setTitle('{{PROJECT_NAME}} API')
    .setDescription('{{PROJECT_NAME}} CRUD API documentation')
    .setVersion('1.0')
    .addTag('{{DOMAIN_LOWER}}')
    .build();
  const document = SwaggerModule.createDocument(app, config);
  SwaggerModule.setup('api', app, document);

  const port = process.env.PORT || 3000;
  await app.listen(port);
  console.log('Application is running on:', port);
}
bootstrap();`, projectName, domainName, gcpProject)

	if err := createFile(filepath.Join(projectName, "src/main.ts"), mainTS); err != nil {
		return err
	}

	// Generate app.module.ts
	appModule := replaceTemplateVars(`import { Module } from '@nestjs/common';
import { MongooseModule } from '@nestjs/mongoose';
import { {{DOMAIN_TITLE}}Module } from './{{DOMAIN_LOWER}}/{{DOMAIN_LOWER}}.module';

@Module({
  imports: [
    MongooseModule.forRoot(process.env.MONGODB_URI || 'mongodb://localhost:27017/{{PROJECT_NAME}}'),
    {{DOMAIN_TITLE}}Module,
  ],
})
export class AppModule {}`, projectName, domainName, gcpProject)

	if err := createFile(filepath.Join(projectName, "src/app.module.ts"), appModule); err != nil {
		return err
	}

	// Generate entity/schema
	schema := replaceTemplateVars(`import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { Document } from 'mongoose';

export type {{DOMAIN_TITLE}}Document = {{DOMAIN_TITLE}} & Document;

@Schema({ timestamps: true })
export class {{DOMAIN_TITLE}} {
  @Prop({ required: true })
  name: string;

  @Prop()
  description?: string;

  @Prop({ default: true })
  isActive: boolean;

  @Prop()
  createdAt?: Date;

  @Prop()
  updatedAt?: Date;
}

export const {{DOMAIN_TITLE}}Schema = SchemaFactory.createForClass({{DOMAIN_TITLE}});`, projectName, domainName, gcpProject)

	if err := createFile(filepath.Join(projectName, fmt.Sprintf("src/%s/entities/%s.entity.ts", domainName, domainName)), schema); err != nil {
		return err
	}

	// Generate DTOs
	createDTO := replaceTemplateVars(`import { IsString, IsOptional, IsBoolean } from 'class-validator';
import { ApiProperty, ApiPropertyOptional } from '@nestjs/swagger';

export class Create{{DOMAIN_TITLE}}Dto {
  @ApiProperty({ description: 'Name of the {{DOMAIN_LOWER}}' })
  @IsString()
  name: string;

  @ApiPropertyOptional({ description: 'Description of the {{DOMAIN_LOWER}}' })
  @IsOptional()
  @IsString()
  description?: string;

  @ApiPropertyOptional({ description: 'Is {{DOMAIN_LOWER}} active', default: true })
  @IsOptional()
  @IsBoolean()
  isActive?: boolean;
}`, projectName, domainName, gcpProject)

	if err := createFile(filepath.Join(projectName, fmt.Sprintf("src/%s/dto/create-%s.dto.ts", domainName, domainName)), createDTO); err != nil {
		return err
	}

	updateDTO := replaceTemplateVars(`import { PartialType } from '@nestjs/swagger';
import { Create{{DOMAIN_TITLE}}Dto } from './create-{{DOMAIN_LOWER}}.dto';

export class Update{{DOMAIN_TITLE}}Dto extends PartialType(Create{{DOMAIN_TITLE}}Dto) {}`, projectName, domainName, gcpProject)

	if err := createFile(filepath.Join(projectName, fmt.Sprintf("src/%s/dto/update-%s.dto.ts", domainName, domainName)), updateDTO); err != nil {
		return err
	}

	// Generate service
	service := replaceTemplateVars(`import { Injectable, NotFoundException } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';
import { {{DOMAIN_TITLE}}, {{DOMAIN_TITLE}}Document } from './entities/{{DOMAIN_LOWER}}.entity';
import { Create{{DOMAIN_TITLE}}Dto } from './dto/create-{{DOMAIN_LOWER}}.dto';
import { Update{{DOMAIN_TITLE}}Dto } from './dto/update-{{DOMAIN_LOWER}}.dto';

@Injectable()
export class {{DOMAIN_TITLE}}Service {
  constructor(
    @InjectModel({{DOMAIN_TITLE}}.name) private {{DOMAIN_LOWER}}Model: Model<{{DOMAIN_TITLE}}Document>,
  ) {}

  async create(create{{DOMAIN_TITLE}}Dto: Create{{DOMAIN_TITLE}}Dto): Promise<{{DOMAIN_TITLE}}> {
    const created{{DOMAIN_TITLE}} = new this.{{DOMAIN_LOWER}}Model(create{{DOMAIN_TITLE}}Dto);
    return created{{DOMAIN_TITLE}}.save();
  }

  async findAll(): Promise<{{DOMAIN_TITLE}}[]> {
    return this.{{DOMAIN_LOWER}}Model.find({ isActive: true }).exec();
  }

  async findOne(id: string): Promise<{{DOMAIN_TITLE}}> {
    const {{DOMAIN_LOWER}} = await this.{{DOMAIN_LOWER}}Model.findById(id).exec();
    if (!{{DOMAIN_LOWER}}) {
      throw new NotFoundException('{{DOMAIN_TITLE}} not found');
    }
    return {{DOMAIN_LOWER}};
  }

  async update(id: string, update{{DOMAIN_TITLE}}Dto: Update{{DOMAIN_TITLE}}Dto): Promise<{{DOMAIN_TITLE}}> {
    const updated{{DOMAIN_TITLE}} = await this.{{DOMAIN_LOWER}}Model
      .findByIdAndUpdate(id, update{{DOMAIN_TITLE}}Dto, { new: true })
      .exec();
    if (!updated{{DOMAIN_TITLE}}) {
      throw new NotFoundException('{{DOMAIN_TITLE}} not found');
    }
    return updated{{DOMAIN_TITLE}};
  }

  async remove(id: string): Promise<void> {
    const result = await this.{{DOMAIN_LOWER}}Model.findByIdAndUpdate(
      id,
      { isActive: false },
      { new: true }
    ).exec();
    if (!result) {
      throw new NotFoundException('{{DOMAIN_TITLE}} not found');
    }
  }
}`, projectName, domainName, gcpProject)

	if err := createFile(filepath.Join(projectName, fmt.Sprintf("src/%s/%s.service.ts", domainName, domainName)), service); err != nil {
		return err
	}

	// Generate controller
	controller := replaceTemplateVars(`import {
  Controller,
  Get,
  Post,
  Body,
  Patch,
  Param,
  Delete,
  HttpCode,
  HttpStatus,
} from '@nestjs/common';
import { ApiTags, ApiOperation, ApiResponse } from '@nestjs/swagger';
import { {{DOMAIN_TITLE}}Service } from './{{DOMAIN_LOWER}}.service';
import { Create{{DOMAIN_TITLE}}Dto } from './dto/create-{{DOMAIN_LOWER}}.dto';
import { Update{{DOMAIN_TITLE}}Dto } from './dto/update-{{DOMAIN_LOWER}}.dto';

@ApiTags('{{DOMAIN_LOWER}}')
@Controller('{{DOMAIN_LOWER}}')
export class {{DOMAIN_TITLE}}Controller {
  constructor(private readonly {{DOMAIN_LOWER}}Service: {{DOMAIN_TITLE}}Service) {}

  @Post()
  @ApiOperation({ summary: 'Create {{DOMAIN_LOWER}}' })
  @ApiResponse({ status: 201, description: '{{DOMAIN_TITLE}} created successfully.' })
  create(@Body() create{{DOMAIN_TITLE}}Dto: Create{{DOMAIN_TITLE}}Dto) {
    return this.{{DOMAIN_LOWER}}Service.create(create{{DOMAIN_TITLE}}Dto);
  }

  @Get()
  @ApiOperation({ summary: 'Get all {{DOMAIN_LOWER}}s' })
  @ApiResponse({ status: 200, description: 'Return all {{DOMAIN_LOWER}}s.' })
  findAll() {
    return this.{{DOMAIN_LOWER}}Service.findAll();
  }

  @Get(':id')
  @ApiOperation({ summary: 'Get {{DOMAIN_LOWER}} by id' })
  @ApiResponse({ status: 200, description: 'Return {{DOMAIN_LOWER}} by id.' })
  @ApiResponse({ status: 404, description: '{{DOMAIN_TITLE}} not found.' })
  findOne(@Param('id') id: string) {
    return this.{{DOMAIN_LOWER}}Service.findOne(id);
  }

  @Patch(':id')
  @ApiOperation({ summary: 'Update {{DOMAIN_LOWER}}' })
  @ApiResponse({ status: 200, description: '{{DOMAIN_TITLE}} updated successfully.' })
  @ApiResponse({ status: 404, description: '{{DOMAIN_TITLE}} not found.' })
  update(@Param('id') id: string, @Body() update{{DOMAIN_TITLE}}Dto: Update{{DOMAIN_TITLE}}Dto) {
    return this.{{DOMAIN_LOWER}}Service.update(id, update{{DOMAIN_TITLE}}Dto);
  }

  @Delete(':id')
  @HttpCode(HttpStatus.NO_CONTENT)
  @ApiOperation({ summary: 'Delete {{DOMAIN_LOWER}}' })
  @ApiResponse({ status: 204, description: '{{DOMAIN_TITLE}} deleted successfully.' })
  @ApiResponse({ status: 404, description: '{{DOMAIN_TITLE}} not found.' })
  remove(@Param('id') id: string) {
    return this.{{DOMAIN_LOWER}}Service.remove(id);
  }
}`, projectName, domainName, gcpProject)

	if err := createFile(filepath.Join(projectName, fmt.Sprintf("src/%s/%s.controller.ts", domainName, domainName)), controller); err != nil {
		return err
	}

	// Generate module
	module := replaceTemplateVars(`import { Module } from '@nestjs/common';
import { MongooseModule } from '@nestjs/mongoose';
import { {{DOMAIN_TITLE}}Service } from './{{DOMAIN_LOWER}}.service';
import { {{DOMAIN_TITLE}}Controller } from './{{DOMAIN_LOWER}}.controller';
import { {{DOMAIN_TITLE}}, {{DOMAIN_TITLE}}Schema } from './entities/{{DOMAIN_LOWER}}.entity';

@Module({
  imports: [
    MongooseModule.forFeature([{ name: {{DOMAIN_TITLE}}.name, schema: {{DOMAIN_TITLE}}Schema }]),
  ],
  controllers: [{{DOMAIN_TITLE}}Controller],
  providers: [{{DOMAIN_TITLE}}Service],
  exports: [{{DOMAIN_TITLE}}Service],
})
export class {{DOMAIN_TITLE}}Module {}`, projectName, domainName, gcpProject)

	if err := createFile(filepath.Join(projectName, fmt.Sprintf("src/%s/%s.module.ts", domainName, domainName)), module); err != nil {
		return err
	}

	// Generate GCP Metrics Interceptor
	interceptor := replaceTemplateVars(`import {
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
    this.projectId = process.env.GCP_PROJECT_ID || '{{GCP_PROJECT}}';
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
      
      // Create custom metric for request duration
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
          type: 'custom.googleapis.com/{{PROJECT_NAME}}/request_duration',
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
            namespace: '{{PROJECT_NAME}}',
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
      const log = this.logging.log('{{PROJECT_NAME}}-requests');
      const metadata = {
        resource: {
          type: 'generic_node',
          labels: {
            location: 'global',
            namespace: '{{PROJECT_NAME}}',
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
}`, projectName, domainName, gcpProject)

	if err := createFile(filepath.Join(projectName, "src/common/interceptors/gcp-metrics.interceptor.ts"), interceptor); err != nil {
		return err
	}

	// Generate Dockerfile
	dockerfile := replaceTemplateVars(`FROM node:18-alpine AS builder

WORKDIR /app

# Copy package files
COPY package*.json ./
RUN npm ci --only=production && npm cache clean --force

FROM node:18-alpine AS development

WORKDIR /app

# Copy package files
COPY package*.json ./
RUN npm ci

# Copy source code
COPY . .

# Build the application
RUN npm run build

FROM node:18-alpine AS production

WORKDIR /app

# Copy production dependencies
COPY --from=builder /app/node_modules ./node_modules

# Copy built application
COPY --from=development /app/dist ./dist

# Create non-root user
RUN addgroup -g 1001 -S nodejs
RUN adduser -S nestjs -u 1001

# Change ownership of the app directory
RUN chown -R nestjs:nodejs /app
USER nestjs

# Expose port
EXPOSE 3000

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD curl -f http://localhost:3000/health || exit 1

# Start the application
CMD ["node", "dist/main.js"]`, projectName, domainName, gcpProject)

	if err := createFile(filepath.Join(projectName, "Dockerfile"), dockerfile); err != nil {
		return err
	}

	// Generate Makefile
	makefile := replaceTemplateVars(`# {{PROJECT_NAME}} Makefile

.PHONY: help install build start dev test lint format clean docker-build docker-run docker-push deploy

# Variables
PROJECT_NAME={{PROJECT_NAME}}
GCP_PROJECT={{GCP_PROJECT}}
IMAGE_NAME=gcr.io/$(GCP_PROJECT)/$(PROJECT_NAME)
VERSION=$(shell git rev-parse --short HEAD)

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

install: ## Install dependencies
	npm install

build: ## Build the application
	npm run build

start: ## Start the application in production mode
	npm run start:prod

dev: ## Start the application in development mode
	npm run start:dev

test: ## Run tests
	npm run test

test-e2e: ## Run e2e tests
	npm run test:e2e

test-cov: ## Run tests with coverage
	npm run test:cov

lint: ## Run linter
	npm run lint

format: ## Format code
	npm run format

clean: ## Clean build artifacts
	rm -rf dist
	rm -rf node_modules
	rm -rf coverage

# Docker commands
docker-build: ## Build Docker image
	docker build -t $(IMAGE_NAME):$(VERSION) .
	docker tag $(IMAGE_NAME):$(VERSION) $(IMAGE_NAME):latest

docker-run: ## Run Docker container locally
	docker run -p 3000:3000 --env-file .env $(IMAGE_NAME):latest

docker-push: ## Push Docker image to GCR
	docker push $(IMAGE_NAME):$(VERSION)
	docker push $(IMAGE_NAME):latest

# GCP commands
gcp-configure: ## Configure GCP CLI
	gcloud config set project $(GCP_PROJECT)
	gcloud auth configure-docker

deploy: docker-build docker-push ## Deploy to Google Cloud Run
	gcloud run deploy $(PROJECT_NAME) \
		--image $(IMAGE_NAME):$(VERSION) \
		--platform managed \
		--region us-central1 \
		--allow-unauthenticated \
		--set-env-vars GCP_PROJECT_ID=$(GCP_PROJECT)

# Development commands
setup: install ## Setup development environment
	@echo "Setting up {{PROJECT_NAME}} development environment..."
	@echo "Creating .env file..."
	@echo "PORT=3000" > .env
	@echo "MONGODB_URI=mongodb://localhost:27017/{{PROJECT_NAME}}" >> .env
	@echo "GCP_PROJECT_ID={{GCP_PROJECT}}" >> .env
	@echo "NODE_ENV=development" >> .env
	@echo "Setup complete!"

logs: ## View application logs
	docker logs -f $(PROJECT_NAME) 2>/dev/null || echo "Container not running"

health: ## Check application health
	curl -f http://localhost:3000/api || echo "Application not responding"`, projectName, domainName, gcpProject)

	if err := createFile(filepath.Join(projectName, "Makefile"), makefile); err != nil {
		return err
	}

	// Generate .env.example
	envExample := replaceTemplateVars(`# Application Configuration
PORT=3000
NODE_ENV=development

# Database Configuration
MONGODB_URI=mongodb://localhost:27017/{{PROJECT_NAME}}

# GCP Configuration
GCP_PROJECT_ID={{GCP_PROJECT}}

# Node Configuration
NODE_NAME={{PROJECT_NAME}}-node`, projectName, domainName, gcpProject)

	if err := createFile(filepath.Join(projectName, ".env.example"), envExample); err != nil {
		return err
	}

	// Generate README.md
	readme := replaceTemplateVars(`# {{PROJECT_NAME}}

{{PROJECT_NAME}} - NestJS CRUD API with GCP integration

## Features

- 🚀 NestJS framework with TypeScript
- 📊 MongoDB integration with Mongoose
- 📈 GCP Monitoring and Logging
- 🐳 Docker support
- 📚 Swagger API documentation
- 🧪 Jest testing setup
- 🔍 ESLint and Prettier
- 🏗️ Clean architecture

## Domain: {{DOMAIN_TITLE}}

This API provides CRUD operations for {{DOMAIN_TITLE}} entities.

## Quick Start

### Prerequisites

- Node.js 18+
- MongoDB
- Docker (optional)
- GCP account (for metrics)

### Installation

1. Clone the repository
2. Install dependencies:
   '''bash
   make install
   '''

3. Setup environment:
   '''bash
   make setup
   '''

4. Update the '.env' file with your configuration

### Running the Application

Development mode:
'''bash
make dev
'''

Production mode:
'''bash
make build
make start
'''

### API Documentation

Once the application is running, visit:
- Swagger UI: http://localhost:3000/api

### API Endpoints

- 'GET /{{DOMAIN_LOWER}}' - Get all {{DOMAIN_LOWER}}s
- 'GET /{{DOMAIN_LOWER}}/:id' - Get {{DOMAIN_LOWER}} by ID
- 'POST /{{DOMAIN_LOWER}}' - Create new {{DOMAIN_LOWER}}
- 'PATCH /{{DOMAIN_LOWER}}/:id' - Update {{DOMAIN_LOWER}}
- 'DELETE /{{DOMAIN_LOWER}}/:id' - Delete {{DOMAIN_LOWER}}

### Testing

Run tests:
'''bash
make test
'''

Run tests with coverage:
'''bash
make test-cov
'''

### Docker

Build and run with Docker:
'''bash
make docker-build
make docker-run
'''

### Deployment

Deploy to Google Cloud Run:
'''bash
make gcp-configure
make deploy
'''

## Project Structure

'''
src/
├── common/
│   └── interceptors/
│       └── gcp-metrics.interceptor.ts
├── {{DOMAIN_LOWER}}/
│   ├── dto/
│   │   ├── create-{{DOMAIN_LOWER}}.dto.ts
│   │   └── update-{{DOMAIN_LOWER}}.dto.ts
│   ├── entities/
│   │   └── {{DOMAIN_LOWER}}.entity.ts
│   ├── {{DOMAIN_LOWER}}.controller.ts
│   ├── {{DOMAIN_LOWER}}.module.ts
│   └── {{DOMAIN_LOWER}}.service.ts
├── app.module.ts
└── main.ts
'''

## Environment Variables

See '.env.example' for all available environment variables.

## GCP Integration

This project includes GCP Monitoring and Logging integration:

- Custom metrics for request duration
- Structured logging
- Error tracking
- Performance monitoring

## License

MIT`, projectName, domainName, gcpProject)

	if err := createFile(filepath.Join(projectName, "README.md"), readme); err != nil {
		return err
	}

	// Generate TypeScript configuration
	tsConfig := `{
  "compilerOptions": {
    "module": "commonjs",
    "declaration": true,
    "removeComments": true,
    "emitDecoratorMetadata": true,
    "experimentalDecorators": true,
    "allowSyntheticDefaultImports": true,
    "target": "ES2021",
    "sourceMap": true,
    "outDir": "./dist",
    "baseUrl": "./",
    "incremental": true,
    "skipLibCheck": true,
    "strictNullChecks": false,
    "noImplicitAny": false,
    "strictBindCallApply": false,
    "forceConsistentCasingInFileNames": false,
    "noFallthroughCasesInSwitch": false
  }
}`

	if err := createFile(filepath.Join(projectName, "tsconfig.json"), tsConfig); err != nil {
		return err
	}

	return nil
}
