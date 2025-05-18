import React, { useState, useEffect } from 'react';
import { Table, Form, Button, Row, Col, Alert } from 'react-bootstrap';
import api from '../api';

export default function ResourceView({ resourceKey, resources }) {
  // найдём конфиг для текущего ресурса
  const findCfg = (key, list) => {
    for (const res of list) {
      if (res.key === key) return res;
      if (res.children) {
        const child = findCfg(key, res.children);
        if (child) return child;
      }
    }
    return null;
  };
  const cfgMeta = findCfg(resourceKey, resources);
  if (!cfgMeta) {
    return <Alert variant="warning">Unknown resource: {resourceKey}</Alert>;
  }

  // определения полей (они неизменились, их можно вынести отдельно,
  // но я оставлю здесь для полноты)
  const meta = {
    sports:             { fields: ['id','name','min_team_size','max_team_size','description'], newFields: ['name','min_team_size','max_team_size','description'] },
    countries:          { fields: ['id','name'], newFields: ['name'] },
    currencies:         { fields: ['code','name'], newFields: ['code','name'] },
    locations:          { fields: ['id','country_id','state','city','address','full_address'], newFields: ['country_id','state','city','address','full_address'] },
    roles:              { fields: ['id','name'], newFields: ['name'] },
    'competition-levels':{ fields: ['id','name'], newFields: ['name'] },
    persons:            { fields: ['id','name','birth_date','country_id'], newFields: ['name','birth_date','country_id'] },
    'persons/sport':    { fields: ['id','person_id','sport_id'], newFields: ['person_id','sport_id'] },
    'persons/team':     { fields: ['id','team_id','person_id','role_id','joined_at','left_at'], newFields: ['team_id','person_id','role_id','joined_at','left_at'] },
    teams:              { fields: ['id','name','country_id','foundation_date'], newFields: ['name','country_id','foundation_date'] },
    competitions:       { fields: ['id','name','sport_id','location_id','level_id'], newFields: ['name','sport_id','location_id','level_id'] },
    stages:             { fields: ['id','competition_id','name','start_time','end_time'], newFields: ['competition_id','name','start_time','end_time'] },
    matches:            { fields: ['id','stage_id','match_time','location_id','metadata'], newFields: ['stage_id','match_time','location_id','metadata'] },
    'match-participants':{ fields: ['id','match_id','team_id','score','is_winner'], newFields: ['match_id','team_id','score','is_winner'] },
    prizes:             { fields: ['id','competition_id','place_bracket','currency_code','prize_amount'], newFields: ['competition_id','place_bracket','currency_code','prize_amount'] },
    'competition-teams':{ fields: ['id','team_id','competition_id'], newFields: ['team_id','competition_id'] },
    'team-achievements':{ fields: ['id','team_id','prize_id'], newFields: ['team_id','prize_id'] },
  };
  const { fields, newFields } = meta[resourceKey];

  const [items, setItems]     = useState([]);
  const [formData, setFormData] = useState(
    newFields.reduce((acc, f) => ({ ...acc, [f]: '' }), {})
  );
  const [loading, setLoading] = useState(true);
  const [error, setError]     = useState(null);

  // загрузка списка
  useEffect(() => {
    setLoading(true);
    setError(null);
    api.get(`/${resourceKey}`)
      .then(res => {
        const data = res.data.data ?? res.data;
        setItems(data);
      })
      .catch(err => setError(err.message))
      .finally(() => setLoading(false));

    // сброс формы
    setFormData(newFields.reduce((acc, f) => ({ ...acc, [f]: '' }), {}));
  }, [resourceKey]);

  const handleChange = f => e =>
    setFormData(prev => ({ ...prev, [f]: e.target.value }));

  const handleSubmit = e => {
    e.preventDefault();
    api.post(`/${resourceKey}`, formData)
      .then(res => {
        setItems(prev => [...prev, res.data]);
        setFormData(newFields.reduce((acc, f) => ({ ...acc, [f]: '' }), {}));
      })
      .catch(err => alert('Error: ' + err.message));
  };

  // состояния
  if (loading) return <div>Loading {cfgMeta.pluralLabel}…</div>;
  if (error)   return <Alert variant="danger">Error: {error}</Alert>;

  return (
    <>
      <h2>{cfgMeta.pluralLabel}</h2>

      <Form onSubmit={handleSubmit} className="mb-4">
        <Row className="g-3">
          {newFields.map(f => (
            <Col key={f} xs={12} sm={6} md={4} lg={3}>
              <Form.Group controlId={`form-${f}`}>
                <Form.Label>{f.replace(/_/g, ' ')}</Form.Label>
                <Form.Control
                  size="sm"
                  type="text"
                  value={formData[f]}
                  onChange={handleChange(f)}
                />
              </Form.Group>
            </Col>
          ))}
        </Row>
        <Button size="sm" className="mt-3">
          Add {cfgMeta.singularLabel}
        </Button>
      </Form>

      <div style={{ overflowX: 'auto' }}>
        <Table size="sm" striped bordered hover>
          <thead>
            <tr>
              {fields.map(f => <th key={f}>{f.replace(/_/g, ' ')}</th>)}
            </tr>
          </thead>
          <tbody>
            {items.map(item => (
              <tr key={item.id ?? JSON.stringify(item)}>
                {fields.map(f => <td key={f}>{String(item[f] ?? '')}</td>)}
              </tr>
            ))}
          </tbody>
        </Table>
      </div>
    </>
  );
}
