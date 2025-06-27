import { Module } from '@nestjs/common';
import { MongooseModule } from '@nestjs/mongoose';
import { {{.DomainTitle}}Service } from './{{.DomainLower}}.service';
import { {{.DomainTitle}}Controller } from './{{.DomainLower}}.controller';
import { {{.DomainTitle}}, {{.DomainTitle}}Schema } from './entities/{{.DomainLower}}.entity';

@Module({
  imports: [
    MongooseModule.forFeature([{ name: {{.DomainTitle}}.name, schema: {{.DomainTitle}}Schema }]),
  ],
  controllers: [{{.DomainTitle}}Controller],
  providers: [{{.DomainTitle}}Service],
  exports: [{{.DomainTitle}}Service],
})
export class {{.DomainTitle}}Module {}
