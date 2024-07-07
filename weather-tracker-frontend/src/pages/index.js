import { useState } from 'react';
import axios from 'axios';
import WeatherCard from '../components/WeatherCard';
import styles from '../styles/Home.module.css';

export default function Home() {
  const [city, setCity] = useState('');
  const [weather, setWeather] = useState(null);
  const [error, setError] = useState('');

  const fetchWeather = async (city) => {
    try {
      const response = await axios.get(`${process.env.NEXT_PUBLIC_BACKEND_URL}/weather/${city}`);
      setWeather(response.data);
      setError('');
    } catch (error) {
      setWeather(null);
      setError('City not found or server error');
    }
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    if (city) {
      fetchWeather(city);
    }
  };

  return (
    <div className={styles.container}>
      <h1>Weather Tracker</h1>
      <form onSubmit={handleSubmit} className={styles.form}>
        <input
          type="text"
          placeholder="Enter city"
          value={city}
          onChange={(e) => setCity(e.target.value)}
          className={styles.input}
        />
        <button type="submit" className={styles.button}>Get Weather</button>
      </form>
      {error && <p className={styles.error}>{error}</p>}
      {weather && <WeatherCard weather={weather} />}
    </div>
  );
}
