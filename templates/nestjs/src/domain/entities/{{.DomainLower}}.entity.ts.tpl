import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { Document } from 'mongoose';

export type {{.DomainTitle}}Document = {{.DomainTitle}} & Document;

@Schema({ timestamps: true })
export class {{.DomainTitle}} {
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

export const {{.DomainTitle}}Schema = SchemaFactory.createForClass({{.DomainTitle}});
