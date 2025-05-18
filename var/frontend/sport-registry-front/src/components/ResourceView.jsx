import React, { useState, useEffect } from 'react';
import { Table, Form, Button, Row, Col, Alert } from 'react-bootstrap';
import api from '../api';

export default function ResourceView({ resourceKey, resources }) {
  // 1. Найти метаданные о ресурсе из списка, чтобы получить labels и children
  const findMeta = (key, list) => {
    for (const r of list) {
      if (r.key === key) return r;
      if (r.children) {
        const c = findMeta(key, r.children);
        if (c) return c;
      }
    }
    return null;
  };
  const cfgMeta = findMeta(resourceKey, resources);
  if (!cfgMeta) {
    return <Alert variant="warning">Unknown resource: {resourceKey}</Alert>;
  }

  // 2. Описать поля и их типы для конвертации перед POST
  const schema = {
    sports:             { name: 'string', min_team_size: 'int', max_team_size: 'int', description: 'string' },
    countries:          { name: 'string' },
    currencies:         { code: 'string', name: 'string' },
    locations:          { country_id: 'int', state: 'string', city: 'string', address: 'string', full_address: 'string' },
    roles:              { name: 'string' },
    'competition-levels':{ name: 'string' },
    persons:            { name: 'string', birth_date: 'string', country_id: 'int' },
    'persons/sport':    { person_id: 'int', sport_id: 'int' },
    'persons/team':     { team_id: 'int', person_id: 'int', role_id: 'int', joined_at: 'string', left_at: 'string' },
    teams:              { name: 'string', country_id: 'int', foundation_date: 'string' },
    competitions:       { name: 'string', sport_id: 'int', location_id: 'int', level_id: 'int' },
    stages:             { competition_id: 'int', name: 'string', start_time: 'string', end_time: 'string' },
    matches:            { stage_id: 'int', match_time: 'string', location_id: 'int', metadata: 'json' },
    'match-participants':{ match_id: 'int', team_id: 'int', score: 'int', is_winner: 'boolean' },
    prizes:             { competition_id: 'int', place_bracket: 'string', currency_code: 'string', prize_amount: 'int' },
    'competition-teams':{ team_id: 'int', competition_id: 'int' },
    'team-achievements':{ team_id: 'int', prize_id: 'int' },
  };

  // 3. ПОЛЯ ДЛЯ ТАБЛИЦЫ И ФОРМЫ
  const fieldMeta = {
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
  const { fields, newFields } = fieldMeta[resourceKey];

  // 4. Хуки
  const [items, setItems]       = useState([]);
  const [formData, setFormData] = useState(
    newFields.reduce((acc, f) => ({ ...acc, [f]: '' }), {})
  );
  const [loading, setLoading] = useState(true);
  const [error, setError]     = useState(null);

  // 5. Загрузка списка
  useEffect(() => {
    setLoading(true);
    api.get(`/${resourceKey}`)
      .then(res => setItems(res.data.data ?? res.data))
      .catch(err => setError(err.message))
      .finally(() => setLoading(false));

    // сброс формы
    setFormData(newFields.reduce((acc, f) => ({ ...acc, [f]: '' }), {}));
  }, [resourceKey]);

  // 6. Обработка изменений
  const handleChange = f => e =>
    setFormData(prev => ({ ...prev, [f]: e.target.value }));

  // 7. Конвертация типов
  const convert = (key, value) => {
    const type = schema[resourceKey][key];
    if (!type || type === 'string') return value;
    if (type === 'int')       return parseInt(value, 10);
    if (type === 'boolean')   return value === 'true';
    if (type === 'json')      {
      try { return JSON.parse(value); }
      catch { return {}; }
    }
    return value;
  };

  // 8. Отправка формы
  const handleSubmit = e => {
    e.preventDefault();

    // соберём полезную нагрузку с правильными типами (как раньше)
    const payload = {};
    for (const f of newFields) {
      payload[f] = convert(f, formData[f]);
    }

    // отправляем POST и после успешного создания — рефрешим список
    api.post(`/${resourceKey}`, payload)
      .then(res => {
        api.post(`/${resourceKey}`, payload)
        .then(() => api.get(`/${resourceKey}`))        // после POST сразу GET
        .then(res => {
          const list = res.data.data ?? res.data;
          setItems(list);
          setFormData(newFields.reduce((acc, f) => ({ ...acc, [f]: '' }), {}));
        })
        .catch(err => {
          console.error(err);
          alert('Error: ' + err.message);
        });
      })
      .catch(err => {
        console.error(err);
        alert('Error: ' + err.message);
      });
  };


  // 9. Рендер
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
                  placeholder={schema[resourceKey][f] === 'json' ? '{"..."}' : ''}
                />
              </Form.Group>
            </Col>
          ))}
        </Row>
        <Button type="submit" size="sm" className="mt-3">
          Add {cfgMeta.singularLabel}
        </Button>
      </Form>

      <div style={{ overflowX: 'auto' }}>
        <Table size="sm" striped bordered hover>
          <thead>
            <tr>{fields.map(f => <th key={f}>{f.replace(/_/g, ' ')}</th>)}</tr>
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