import React, { useEffect, useState } from 'react';
import { Bar } from 'react-chartjs-2';
import { Spinner, Alert } from 'react-bootstrap';
import api from '../api';
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  BarElement,
  Title,
  Tooltip,
  Legend
} from 'chart.js';

ChartJS.register(CategoryScale, LinearScale, BarElement, Title, Tooltip, Legend);

export default function Dashboard({ resources }) {
  const [counts, setCounts] = useState({});
  const [loading, setLoading] = useState(true);
  const [error, setError]   = useState(null);

  useEffect(() => {
    setLoading(true);
    Promise.all(
      resources.map(r =>
        api.get(`/${r.key}`)
          .then(res => [r.pluralLabel, (res.data.data ?? res.data).length])
      )
    )
      .then(results => {
        setCounts(Object.fromEntries(results));
      })
      .catch(err => setError(err.message))
      .finally(() => setLoading(false));
  }, [resources]);

  if (loading) return <Spinner animation="border" />;
  if (error)   return <Alert variant="danger">{error}</Alert>;

  const labels = Object.keys(counts);
  const data = {
    labels,
    datasets: [
      {
        label: 'Record count',
        data: labels.map(l => counts[l]),
        backgroundColor: 'rgba(75, 192, 192, 0.5)'
      }
    ]
  };

  return (
    <>
      <h2>Dashboard</h2>
      <Bar data={data} />
    </>
  );
}
