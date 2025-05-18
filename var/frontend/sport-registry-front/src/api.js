import axios from 'axios';

const api = axios.create({
  baseURL: '',
  headers: { 'Sport-Registry-ApiKey': '123456' },
});

export default api;
