import { Module } from '@nestjs/common';
import { OrdersService } from './orders.service';
import { OrdersController } from './orders.controller';
import { SequelizeModule } from '@nestjs/sequelize';
import { Order } from './entities/order.entity';
import { AccountModule } from 'src/account/account.module';

@Module({
  imports: [SequelizeModule.forFeature([Order]), AccountModule],
  controllers: [OrdersController],
  providers: [OrdersService],
})
export class OrdersModule {}
