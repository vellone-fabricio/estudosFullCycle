import { Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/sequelize';
import { EmptyResultError } from 'sequelize';
import { AccountStorageService } from 'src/account/account-storage/account-storage.service';
import { CreateOrderDto } from './dto/create-order.dto';
import { UpdateOrderDto } from './dto/update-order.dto';
import { Order } from './entities/order.entity';

@Injectable()
export class OrdersService {
  constructor(
    @InjectModel(Order)
    private orderModel: typeof Order,
    private accountStorage: AccountStorageService,
  ) {}

  async create(createOrderDto: CreateOrderDto) {
    return await this.orderModel.create({
      ...createOrderDto,
      account_id: this.accountStorage.account.id,
    });
  }

  findAll() {
    return this.orderModel.findAll({
      where: {
        account_id: this.accountStorage.account.id,
      },
    });
  }

  async findOne(id: string) {
    return await this.orderModel.findOne({
      where: {
        id,
        account_id: this.accountStorage.account.id,
      },
      rejectOnEmpty: new EmptyResultError(`Account with ID ${id} not found!`),
    });
  }

  async update(id: string, updateOrderDto: UpdateOrderDto) {
    const order: Order = await this.findOne(id);
    order.update(updateOrderDto);
    return order;
  }

  async remove(id: string) {
    const order: Order = await this.findOne(id);
    return order.destroy();
  }
}
