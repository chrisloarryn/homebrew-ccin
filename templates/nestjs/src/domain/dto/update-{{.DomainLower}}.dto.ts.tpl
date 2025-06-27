import { PartialType } from '@nestjs/swagger';
import { Create{{.DomainTitle}}Dto } from './create-{{.DomainLower}}.dto';

export class Update{{.DomainTitle}}Dto extends PartialType(Create{{.DomainTitle}}Dto) {}
