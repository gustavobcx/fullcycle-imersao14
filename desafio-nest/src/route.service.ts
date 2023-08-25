import { Injectable } from '@nestjs/common';
import { PrismaService } from './prisma.service';
import { Prisma, Route } from '@prisma/client';

@Injectable()
export class RouteService {
  constructor(private prisma: PrismaService) {}

  async find(): Promise<Route[]> {
    return this.prisma.route.findMany();
  }

  async create(data: Prisma.RouteCreateInput): Promise<Route> {
    return this.prisma.route.create({
      data,
    });
  }
}
