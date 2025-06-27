import {
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
import { {{.DomainTitle}}Service } from './{{.DomainLower}}.service';
import { Create{{.DomainTitle}}Dto } from './dto/create-{{.DomainLower}}.dto';
import { Update{{.DomainTitle}}Dto } from './dto/update-{{.DomainLower}}.dto';

@ApiTags('{{.DomainLower}}')
@Controller('{{.DomainLower}}')
export class {{.DomainTitle}}Controller {
  constructor(private readonly {{.DomainLower}}Service: {{.DomainTitle}}Service) {}

  @Post()
  @ApiOperation({ summary: 'Create {{.DomainLower}}' })
  @ApiResponse({ status: 201, description: '{{.DomainTitle}} created successfully.' })
  create(@Body() create{{.DomainTitle}}Dto: Create{{.DomainTitle}}Dto) {
    return this.{{.DomainLower}}Service.create(create{{.DomainTitle}}Dto);
  }

  @Get()
  @ApiOperation({ summary: 'Get all {{.DomainLower}}s' })
  @ApiResponse({ status: 200, description: 'Return all {{.DomainLower}}s.' })
  findAll() {
    return this.{{.DomainLower}}Service.findAll();
  }

  @Get(':id')
  @ApiOperation({ summary: 'Get {{.DomainLower}} by id' })
  @ApiResponse({ status: 200, description: 'Return {{.DomainLower}} by id.' })
  @ApiResponse({ status: 404, description: '{{.DomainTitle}} not found.' })
  findOne(@Param('id') id: string) {
    return this.{{.DomainLower}}Service.findOne(id);
  }

  @Patch(':id')
  @ApiOperation({ summary: 'Update {{.DomainLower}}' })
  @ApiResponse({ status: 200, description: '{{.DomainTitle}} updated successfully.' })
  @ApiResponse({ status: 404, description: '{{.DomainTitle}} not found.' })
  update(@Param('id') id: string, @Body() update{{.DomainTitle}}Dto: Update{{.DomainTitle}}Dto) {
    return this.{{.DomainLower}}Service.update(id, update{{.DomainTitle}}Dto);
  }

  @Delete(':id')
  @HttpCode(HttpStatus.NO_CONTENT)
  @ApiOperation({ summary: 'Delete {{.DomainLower}}' })
  @ApiResponse({ status: 204, description: '{{.DomainTitle}} deleted successfully.' })
  @ApiResponse({ status: 404, description: '{{.DomainTitle}} not found.' })
  remove(@Param('id') id: string) {
    return this.{{.DomainLower}}Service.remove(id);
  }
}
