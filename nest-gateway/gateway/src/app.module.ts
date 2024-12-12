import { Module } from '@nestjs/common';
import { HttpModule } from '@nestjs/axios';
import { PythonController } from './python/python.controller';
import { GoController } from './go/go.controller';

@Module({
  imports: [
    HttpModule.register({
      timeout: 5000,
      maxRedirects: 5,
    }),
  ],
  controllers: [PythonController, GoController],
})
export class AppModule {}

