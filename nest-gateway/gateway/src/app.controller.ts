import { Controller, Get, Post, Put, Delete, Body, Param } from '@nestjs/common';
import { HttpService } from '@nestjs/axios';
import { firstValueFrom } from 'rxjs';

@Controller('python')
export class PythonController {
  constructor(private httpService: HttpService) {}

  private readonly baseUrl = 'http://localhost:8082/api';

  // User endpoints
  @Get('users')
  async getAllUsers() {
    try {
      const { data } = await firstValueFrom(this.httpService.get(`${this.baseUrl}/users`));
      return data;
    } catch (error) {
      console.error('Error fetching users:', error);
      throw error;
    }
  }

  @Get('users/:id')
  async getUserById(@Param('id') id: string) {
    try {
      const { data } = await firstValueFrom(this.httpService.get(`${this.baseUrl}/users/${id}`));
      return data;
    } catch (error) {
      console.error(`Error fetching user ${id}:`, error);
      throw error;
    }
  }

  @Post('users')
  async createUser(@Body() createUserDto: any) {
    try {
      const { data } = await firstValueFrom(this.httpService.post(`${this.baseUrl}/users`, createUserDto));
      return data;
    } catch (error) {
      console.error('Error creating user:', error);
      throw error;
    }
  }

  @Put('users/:id')
  async updateUser(@Param('id') id: string, @Body() updateUserDto: any) {
    try {
      const { data } = await firstValueFrom(this.httpService.put(`${this.baseUrl}/users/${id}`, updateUserDto));
      return data;
    } catch (error) {
      console.error(`Error updating user ${id}:`, error);
      throw error;
    }
  }

  @Delete('users/:id')
  async deleteUser(@Param('id') id: string) {
    try {
      const { data } = await firstValueFrom(this.httpService.delete(`${this.baseUrl}/users/${id}`));
      return data;
    } catch (error) {
      console.error(`Error deleting user ${id}:`, error);
      throw error;
    }
  }

  // Company endpoints
  @Get('companies')
  async getAllCompanies() {
    try {
      const { data } = await firstValueFrom(this.httpService.get(`${this.baseUrl}/companies`));
      return data;
    } catch (error) {
      console.error('Error fetching companies:', error);
      throw error;
    }
  }

  @Get('companies/:id')
  async getCompanyById(@Param('id') id: string) {
    try {
      const { data } = await firstValueFrom(this.httpService.get(`${this.baseUrl}/companies/${id}`));
      return data;
    } catch (error) {
      console.error(`Error fetching company ${id}:`, error);
      throw error;
    }
  }

  @Post('companies')
  async createCompany(@Body() createCompanyDto: any) {
    try {
      const { data } = await firstValueFrom(this.httpService.post(`${this.baseUrl}/companies`, createCompanyDto));
      return data;
    } catch (error) {
      console.error('Error creating company:', error);
      throw error;
    }
  }

  @Put('companies/:id')
  async updateCompany(@Param('id') id: string, @Body() updateCompanyDto: any) {
    try {
      const { data } = await firstValueFrom(this.httpService.put(`${this.baseUrl}/companies/${id}`, updateCompanyDto));
      return data;
    } catch (error) {
      console.error(`Error updating company ${id}:`, error);
      throw error;
    }
  }

  @Delete('companies/:id')
  async deleteCompany(@Param('id') id: string) {
    try {
      const { data } = await firstValueFrom(this.httpService.delete(`${this.baseUrl}/companies/${id}`));
      return data;
    } catch (error) {
      console.error(`Error deleting company ${id}:`, error);
      throw error;
    }
  }
}

