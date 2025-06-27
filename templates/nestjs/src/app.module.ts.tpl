import { Module } from '@nestjs/common';
import { MongooseModule } from '@nestjs/mongoose';
import { {{.DomainTitle}}Module } from './{{.DomainLower}}/{{.DomainLower}}.module';

@Module({
  imports: [
    MongooseModule.forRoot(process.env.MONGODB_URI || '{{.DatabaseType}}://localhost:27017/{{.ProjectName}}'),
    {{.DomainTitle}}Module,
  ],
})
export class AppModule {}
