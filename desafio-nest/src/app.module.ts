import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { PrismaService } from './prisma.service';
import { RouteService } from './route.service';

@Module({
  imports: [],
  controllers: [AppController],
  providers: [PrismaService, RouteService],
})
export class AppModule {}
