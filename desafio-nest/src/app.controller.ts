import { Controller, Get, Post, Body } from '@nestjs/common';
import { RouteService } from './route.service';
import { Coordinate, Route as RouteModel } from '@prisma/client';

@Controller('api')
export class AppController {
  constructor(private readonly routeService: RouteService) {}

  @Get('routes')
  async find(): Promise<RouteModel[]> {
    return this.routeService.find();
  }

  @Post('routes')
  async create(
    @Body()
    routeData: {
      name: string;
      source: Coordinate;
      destination: Coordinate;
    },
  ): Promise<RouteModel> {
    const { name, source, destination } = routeData;
    return this.routeService.create({
      name,
      source,
      destination,
    });
  }
}
