
import axios from 'axios';

const API_BASE_URL = 'http://localhost:8080';

export const api = {
  async getMonitoringData() {
    const response = await axios.get(`${API_BASE_URL}/monitoring`);
    return response.data;
  },

  async getAllMapPlaces() {
    const response = await axios.get(`${API_BASE_URL}/map-places`);
    return response.data;
  },

  async getAllTransactions() {
    const response = await axios.get(`${API_BASE_URL}/transactions`);
    return response.data;
  }
};