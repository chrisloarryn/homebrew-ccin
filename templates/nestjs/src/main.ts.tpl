import { NestFactory } from '@nestjs/core';
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
    .setTitle('{{.ProjectName}} API')
    .setDescription('{{.ProjectName}} CRUD API documentation')
    .setVersion('1.0')
    .addTag('{{.DomainLower}}')
    .build();
  const document = SwaggerModule.createDocument(app, config);
  SwaggerModule.setup('api', app, document);

  const port = process.env.PORT || {{.Port}};
  await app.listen(port);
  console.log(`Application is running on: ${port}`);
}
bootstrap();
