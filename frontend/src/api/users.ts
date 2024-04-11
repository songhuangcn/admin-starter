import axios from 'axios';

export interface User {
  id: number;
  username: string;
  name: string;
}

export function usersIndex(params: any) {
  return axios.get('/api/users', { params });
}

export function usersCreate(data: any) {
  return axios.post('/api/users', data);
}

export function usersUpdate(id: any, data: any) {
  return axios.put(`/api/users/${id}`, data);
}

export function usersDestroy(id: any) {
  return axios.delete(`/api/users/${id}`);
}
