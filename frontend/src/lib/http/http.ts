import axios, { AxiosResponse } from 'axios';
import idsrvAuth from '@/lib/auth/auth';

const defaultOptions = {
  baseURL: `${process.env.VUE_APP_COOKBOOK_BACKEND_URL}/v1`,
  timeout: 15000,
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

export interface IngredientInput {
  ID: number;
  amount: number;
  unit: string;
}

export interface Ingredient {
  id: number;
  name: string;
}

export interface Section {
  id: number;
  name: string;
}

export interface RecipeInput {
  ID: number;
  Title: string;
  Description: string;
  Ingredients: Array<IngredientInput>;
  Method: string;
  PrepTime: number;
  CookTime: number;
  AmountPersons: number;
}

export interface IngredientAmounts {
  recipeid: number;
  ingredientid: number;
  sectionid: number;
  amount: number;
  unit: string;
}

export interface RecipeResponse {
  id: number;
  title: string;
  description: string;
  ingredients: Array<Ingredient>;
  sections: Array<Section>;
  method: string;
  preptime: number;
  cooktime: number;
  persons: number;
  ingredientamounts: Array<IngredientAmounts>;
}

const responseBody = (response: AxiosResponse) => response.data;

const requests = {
  get: (url: string) => instance.get(url).then(responseBody),
  post: (url: string, body: {}) => instance.post(url, body).then(responseBody),
  put: (url: string, body: {}) => instance.put(url, body).then(responseBody),
  delete: (url: string) => instance.delete(url).then(responseBody),
};

export const Recipes = {
  getAllRecipes: (): Promise<RecipeResponse[]> => requests.get('/recipe'),
  getSingleRecipe: (id: number): Promise<RecipeResponse> => requests.get(`/recipe/${id}`),
  createRecipe: (item: RecipeInput): Promise<AxiosResponse> => requests.post('/recipe', JSON.stringify(item)),
};

export const Ingredients = {
  getAllIngredients: (): Promise<Ingredient[]> => requests.get('/ingredients'),
  getSingleIngredient: (ingredientID: number): Promise<Ingredient[]> => requests.get(`/ingredients/${ingredientID}`),
  createIngredient: (item: IngredientInput): Promise<AxiosResponse> => requests.post('/ingredients', JSON.stringify(item)),
};

export const Sections = {
  getAllIngredientSections: (): Promise<Section[]> => requests.get('/sections'),
  getSingleIngredientSection: (sectionID: number): Promise<Section[]> => requests.get(`/sections/${sectionID}`),
};
