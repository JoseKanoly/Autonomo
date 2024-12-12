import { Controller, Get, Post, Put, Delete, Body, Param } from '@nestjs/common';
import { HttpService } from '@nestjs/axios';
import { firstValueFrom } from 'rxjs';

@Controller('go')
export class GoController {
  constructor(private httpService: HttpService) {}

  private readonly baseUrl = 'http://localhost:8081/api'; // Assuming Go service runs on port 8081

  // Lista Precios endpoints
  @Get('listas_precios')
  async getAllListaPrecios() {
    try {
      const { data } = await firstValueFrom(this.httpService.get(`${this.baseUrl}/listas_precios`));
      return data;
    } catch (error) {
      console.error('Error fetching lista precios:', error);
      throw error;
    }
  }

  @Get('listas_precios/:id')
  async getListaPreciosById(@Param('id') id: string) {
    try {
      const { data } = await firstValueFrom(this.httpService.get(`${this.baseUrl}/listas_precios/${id}`));
      return data;
    } catch (error) {
      console.error(`Error fetching lista precios ${id}:`, error);
      throw error;
    }
  }

  @Post('listas_precios')
  async createListaPrecios(@Body() createListaPreciosDto: any) {
    try {
      const { data } = await firstValueFrom(this.httpService.post(`${this.baseUrl}/listas_precios`, createListaPreciosDto));
      return data;
    } catch (error) {
      console.error('Error creating lista precios:', error);
      throw error;
    }
  }

  @Put('listas_precios/:id')
  async updateListaPrecios(@Param('id') id: string, @Body() updateListaPreciosDto: any) {
    try {
      const { data } = await firstValueFrom(this.httpService.put(`${this.baseUrl}/listas_precios/${id}`, updateListaPreciosDto));
      return data;
    } catch (error) {
      console.error(`Error updating lista precios ${id}:`, error);
      throw error;
    }
  }

  @Delete('listas_precios/:id')
  async deleteListaPrecios(@Param('id') id: string) {
    try {
      const { data } = await firstValueFrom(this.httpService.delete(`${this.baseUrl}/listas_precios/${id}`));
      return data;
    } catch (error) {
      console.error(`Error deleting lista precios ${id}:`, error);
      throw error;
    }
  }

  // Residuos endpoints
  @Get('residuos')
  async getAllResiduos() {
    try {
      const { data } = await firstValueFrom(this.httpService.get(`${this.baseUrl}/residuos`));
      return data;
    } catch (error) {
      console.error('Error fetching residuos:', error);
      throw error;
    }
  }

  @Get('residuos/:id')
  async getResiduoById(@Param('id') id: string) {
    try {
      const { data } = await firstValueFrom(this.httpService.get(`${this.baseUrl}/residuos/${id}`));
      return data;
    } catch (error) {
      console.error(`Error fetching residuo ${id}:`, error);
      throw error;
    }
  }

  @Post('residuos')
  async createResiduo(@Body() createResiduoDto: any) {
    try {
      const { data } = await firstValueFrom(this.httpService.post(`${this.baseUrl}/residuos`, createResiduoDto));
      return data;
    } catch (error) {
      console.error('Error creating residuo:', error);
      throw error;
    }
  }

  @Put('residuos/:id')
  async updateResiduo(@Param('id') id: string, @Body() updateResiduoDto: any) {
    try {
      const { data } = await firstValueFrom(this.httpService.put(`${this.baseUrl}/residuos/${id}`, updateResiduoDto));
      return data;
    } catch (error) {
      console.error(`Error updating residuo ${id}:`, error);
      throw error;
    }
  }

  @Delete('residuos/:id')
  async deleteResiduo(@Param('id') id: string) {
    try {
      const { data } = await firstValueFrom(this.httpService.delete(`${this.baseUrl}/residuos/${id}`));
      return data;
    } catch (error) {
      console.error(`Error deleting residuo ${id}:`, error);
      throw error;
    }
  }
}

