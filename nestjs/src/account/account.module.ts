import { Module } from '@nestjs/common';
import { AccountService } from './account.service';
import { AccountController } from './account.controller';
import { SequelizeModule } from '@nestjs/sequelize';
import { Account } from './entities/account.entity';
import { AccountStorageService } from './account-storage/account-storage.service';
import { TokenGuard } from './token.guard';

@Module({
  imports: [SequelizeModule.forFeature([Account])],
  controllers: [AccountController],
  providers: [AccountService, AccountStorageService, TokenGuard],
  exports: [AccountStorageService],
})
export class AccountModule {}
