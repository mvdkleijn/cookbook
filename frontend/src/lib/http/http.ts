import axios, { AxiosResponse } from 'axios';
import idsrvAuth from '@/lib/auth/auth';

const defaultOptions = {
  baseURL: 'http://localhost:8080/v1',
  headers: {
    'Content-Type': 'application/json',
  },
};

const instance = axios.create(defaultOptions);

instance.interceptors.request.use((config) => {
  const conf = config;
  const token = idsrvAuth.accessToken;
  conf.headers.Authorization = token ? `Bearer ${token}` : '';
  return conf;
});

interface RecipeInput {
  ID: number;
  Title: string;
  Description: string;
  Ingredients: Array<IngredientInput>;
  Method: string;
  PrepTime: number;
  CookTime: number;
  TotalTime: number;
  Amount_Persons: number;
}

interface RecipeResponse {
  id: number;
  title: string;
  description: string;
  ingredients: Array<IngredientResponse>;
  method: string;
  preptime: number;
  cooktime: number;
  totaltime: number;
  persons: number;
}

interface IngredientInput {
  ID: number;
  amount: number;
  unit: string;
}

interface IngredientResponse {
  id: number;
  name: string;
}

export function getRecipes(): Promise<AxiosResponse> {
  const response = instance.get('/recipe');
  return response;
}

export function createRecipe(item: any): Promise<AxiosResponse> {
  const response = instance.post('/recipe', JSON.stringify(item));
  return response;
}

export function getIngredients(): Promise<AxiosResponse> {
  const response = instance.get('/ingredient');
  return response;
}

export function createIngredient(item: any): Promise<AxiosResponse> {
  const response = instance.post('/ingredient', JSON.stringify(item));
  return response;
}
