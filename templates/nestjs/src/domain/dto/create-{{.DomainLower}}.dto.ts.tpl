import { IsString, IsOptional, IsBoolean } from 'class-validator';
import { ApiProperty, ApiPropertyOptional } from '@nestjs/swagger';

export class Create{{.DomainTitle}}Dto {
  @ApiProperty({ description: 'Name of the {{.DomainLower}}' })
  @IsString()
  name: string;

  @ApiPropertyOptional({ description: 'Description of the {{.DomainLower}}' })
  @IsOptional()
  @IsString()
  description?: string;

  @ApiPropertyOptional({ description: 'Is {{.DomainLower}} active', default: true })
  @IsOptional()
  @IsBoolean()
  isActive?: boolean;
}
