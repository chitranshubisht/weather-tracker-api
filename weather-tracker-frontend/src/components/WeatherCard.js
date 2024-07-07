import styles from '../styles/WeatherCard.module.css';

export default function WeatherCard({ weather }) {
  return (
    <div className={styles.card}>
      <h2>{weather.name}</h2>
      <p>Temperature in °C : {weather.main.celsius.toFixed(2)} °C</p>
      <p>Temperature in °F: {weather.main.fahrenheit.toFixed(2)} °F</p>
      <p>Temperature °K: {weather.main.temp.toFixed(2)} °K</p>
    </div>
  );
}
