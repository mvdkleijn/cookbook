import axios from 'axios';
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

export function getRecipes(): Promise<any> {
  const response = instance.get('/recipe');
  return response;
}

export function createRecipe(item: any): Promise<any> {
  const response = instance.post('/recipe', JSON.stringify(item));
  return response;
}

export function getIngredients(): Promise<any> {
  const response = instance.get('/ingredient');
  return response;
}

export function createIngredient(item: any): Promise<any> {
  const response = instance.post('/ingredient', JSON.stringify(item));
  return response;
}
