import axios from 'axios';

export const fetchWeather = async (city) => {
  const response = await axios.get(`/weather/${city}`);
  return response.data;
};
