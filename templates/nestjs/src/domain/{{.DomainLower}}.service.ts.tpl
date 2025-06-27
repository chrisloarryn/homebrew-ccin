import { Injectable, NotFoundException } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';
import { {{.DomainTitle}}, {{.DomainTitle}}Document } from './entities/{{.DomainLower}}.entity';
import { Create{{.DomainTitle}}Dto } from './dto/create-{{.DomainLower}}.dto';
import { Update{{.DomainTitle}}Dto } from './dto/update-{{.DomainLower}}.dto';

@Injectable()
export class {{.DomainTitle}}Service {
  constructor(
    @InjectModel({{.DomainTitle}}.name) private {{.DomainLower}}Model: Model<{{.DomainTitle}}Document>,
  ) {}

  async create(create{{.DomainTitle}}Dto: Create{{.DomainTitle}}Dto): Promise<{{.DomainTitle}}> {
    const created{{.DomainTitle}} = new this.{{.DomainLower}}Model(create{{.DomainTitle}}Dto);
    return created{{.DomainTitle}}.save();
  }

  async findAll(): Promise<{{.DomainTitle}}[]> {
    return this.{{.DomainLower}}Model.find({ isActive: true }).exec();
  }

  async findOne(id: string): Promise<{{.DomainTitle}}> {
    const {{.DomainLower}} = await this.{{.DomainLower}}Model.findById(id).exec();
    if (!{{.DomainLower}}) {
      throw new NotFoundException('{{.DomainTitle}} not found');
    }
    return {{.DomainLower}};
  }

  async update(id: string, update{{.DomainTitle}}Dto: Update{{.DomainTitle}}Dto): Promise<{{.DomainTitle}}> {
    const updated{{.DomainTitle}} = await this.{{.DomainLower}}Model
      .findByIdAndUpdate(id, update{{.DomainTitle}}Dto, { new: true })
      .exec();
    if (!updated{{.DomainTitle}}) {
      throw new NotFoundException('{{.DomainTitle}} not found');
    }
    return updated{{.DomainTitle}};
  }

  async remove(id: string): Promise<void> {
    const result = await this.{{.DomainLower}}Model.findByIdAndUpdate(
      id,
      { isActive: false },
      { new: true }
    ).exec();
    if (!result) {
      throw new NotFoundException('{{.DomainTitle}} not found');
    }
  }
}
