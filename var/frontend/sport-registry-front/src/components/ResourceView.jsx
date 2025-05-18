// src/components/ResourceView.js
import React, { useState, useEffect } from 'react';
import {
  Form,
  Button,
  Row,
  Col,
  Alert,
  Modal,
  Spinner,
  Pagination
} from 'react-bootstrap';
import api from '../api';

export default function ResourceView({ resourceKey, resources }) {
  // 1. Найти метаданные (labels, children)
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
  const cfg = findMeta(resourceKey, resources);
  if (!cfg) return <Alert variant="warning">Unknown resource: {resourceKey}</Alert>;

  // 2. Типы полей для placeholder и конвертации
  const schema = {
    sports:              { name: 'string', min_team_size: 'int', max_team_size: 'int', description: 'string' },
    countries:           { name: 'string' },
    currencies:          { code: 'string', name: 'string' },
    locations:           { country_id: 'int', state: 'string', city: 'string', address: 'string', full_address: 'string' },
    roles:               { name: 'string' },
    'competition-teams': { team_id: 'int', competition_id: 'int' },
    'team-achievements': { team_id: 'int', prize_id: 'int' },
    'match-participants':{ match_id: 'int', team_id: 'int', score: 'int', is_winner: 'boolean' },
    persons:             { name: 'string', birth_date: 'string', country_id: 'int' },
    'persons/sport':     { person_id: 'int', sport_id: 'int' },
    'persons/team':      { team_id: 'int', person_id: 'int', role_id: 'int', joined_at: 'string', left_at: 'string' },
    teams:               { name: 'string', country_id: 'int', foundation_date: 'string' },
    competitions:        { name: 'string', sport_id: 'int', location_id: 'int', level_id: 'int' },
    'competition-levels':{ name: 'string' },
    stages:              { competition_id: 'int', name: 'string', start_time: 'string', end_time: 'string' },
    matches:             { stage_id: 'int', match_time: 'string', location_id: 'int', metadata: 'json' },
    prizes:              { competition_id: 'int', place_bracket: 'string', currency_code: 'string', prize_amount: 'int' },
  };

  // 3. Параметры списка и формы
  const fieldMeta = {
    sports:               { labelField: 'name',           newFields: ['name','min_team_size','max_team_size','description'] },
    countries:            { labelField: 'name',           newFields: ['name'] },
    currencies:           { labelField: 'name',           newFields: ['code','name'] },
    locations:            { labelField: 'full_address',   newFields: ['country_id','state','city','address','full_address'] },
    roles:                { labelField: 'name',           newFields: ['name'] },
    'competition-teams':  { labelField: 'team_id',        newFields: ['team_id','competition_id'] },
    'team-achievements':  { labelField: 'team_id',        newFields: ['team_id','prize_id'] },
    'match-participants': { labelField: 'match_id',       newFields: ['match_id','team_id','score','is_winner'] },
    persons:              { labelField: 'name',           newFields: ['name','birth_date','country_id'] },
    'persons/sport':      { labelField: 'sport_id',       newFields: ['person_id','sport_id'] },
    'persons/team':       { labelField: 'person_id',      newFields: ['team_id','person_id','role_id','joined_at','left_at'] },
    teams:                { labelField: 'name',           newFields: ['name','country_id','foundation_date'] },
    competitions:         { labelField: 'name',           newFields: ['name','sport_id','location_id','level_id'] },
    'competition-levels': { labelField: 'name',           newFields: ['name'] },
    stages:               { labelField: 'name',           newFields: ['competition_id','name','start_time','end_time'] },
    matches:              { labelField: 'match_time',     newFields: ['stage_id','match_time','location_id','metadata'] },
    prizes:               { labelField: 'place_bracket',  newFields: ['competition_id','place_bracket','currency_code','prize_amount'] },
  };
  const { labelField, newFields } = fieldMeta[resourceKey];

  // 4. Внешние ключи и лейблы
  const fkMap = {
    country_id: 'countries',
    sport_id:   'sports',
    person_id:  'persons',
    team_id:    'teams',
    role_id:    'roles',
    competition_id: 'competitions',
    stage_id:   'stages',
    match_id:   'matches',
    prize_id:   'prizes',
    currency_code:'currencies',
  };
  const labelMap = {
    countries: 'name', sports: 'name', persons: 'name', teams: 'name',
    roles: 'name', competitions: 'name', stages: 'name', matches: 'match_time',
    prizes: 'place_bracket', locations: 'full_address', currencies: 'name',
  };

  // 5. Ключ сущности
  const primaryKeyMap = { currencies: 'code', default: 'id' };
  const idKey = primaryKeyMap[resourceKey] || primaryKeyMap.default;

  // 6. Состояния
  const [list, setList]               = useState([]);
  const [loading, setLoading]         = useState(true);
  const [error, setError]             = useState(null);
  const [page, setPage]               = useState(1);
  const perPage = 10;

  const [showCreate, setShowCreate]   = useState(false);
  const [formData, setFormData]       = useState(newFields.reduce((a,f)=>(a[f]='',a),{}));
  const [errors, setErrors]           = useState({});
  const [fkOpts, setFkOpts]           = useState({});

  const [showDetail, setShowDetail]   = useState(false);
  const [detailId, setDetailId]       = useState(null);
  const [detail, setDetail]           = useState(null);
  const [detailLoading, setDetailLoading] = useState(false);
  const [detailError, setDetailError]     = useState(null);

  // 7. Генерация placeholder
  const getPlaceholder = f => {
    const t = schema[resourceKey][f];
    if (t === 'int')     return 'e.g. 123';
    if (t === 'boolean') return 'true or false';
    if (t === 'json')    return '{"key":"value"}';
    return `Enter ${f.replace(/_/g, ' ')}`;
  };

  const fetchList = () => {
    setError(null);
    setLoading(true);
    api.get(`/${resourceKey}`)
      .then(res => {
        const data = res.data.data ?? res.data;
        setList(data);
        setPage(1);           // вернём пагинацию на первую страницу
      })
      .catch(err => setError(err.message))
      .finally(() => setLoading(false));
  };

  // 8. Загрузка списка + FK-опций
  useEffect(() => {
    fetchList();

    newFields.forEach(f => {
      const ep = fkMap[f];
      if (!ep) return;
      api.get(`/${ep}`)
        .then(r => setFkOpts(prev => ({ ...prev, [f]: r.data.data ?? r.data })))
        .catch(console.error);
    });

    // сброс формы и ошибок
    setFormData(newFields.reduce((a, f) => ({ ...a, [f]: '' }), {}));
    setErrors({});
  }, [resourceKey]);

  // 9. Конвертация перед отправкой
  const convert = (k, v) => {
    const t = schema[resourceKey][k];
    if (t === 'int')     return parseInt(v, 10);
    if (t === 'boolean') return v === 'true';
    if (t === 'json') {
      try { return JSON.parse(v); }
      catch { return {}; }
    }
    return v;
  };

  // 10. Валидация формы
  const validate = () => {
    const err = {};
    newFields.forEach(f => {
      if (!formData[f]) err[f] = 'Required';
    });
    setErrors(err);
    return Object.keys(err).length === 0;
  };

  // 11. Submit Create
  const handleCreate = e => {
    e.preventDefault();
    if (!validate()) return;

    const payload = {};
    newFields.forEach(f => { payload[f] = convert(f, formData[f]); });

    api.post(`/${resourceKey}`, payload)
      .then(() => {
        fetchList();
        setFormData(newFields.reduce((acc, f) => ({ ...acc, [f]: '' }), {}));
        setErrors({});
        setShowCreate(false);
      })
      .catch(err => alert('Create error: ' + err.message));
  };

  // 12. Открыть детали
  const openDetail = id => {
    setDetailId(id);
    setShowDetail(true);
    setDetail(null);
    setDetailError(null);
    setDetailLoading(true);
    api.get(`/${resourceKey}/${id}`)
      .then(r => setDetail(r.data.data ?? r.data))
      .catch(e => setDetailError(e.message))
      .finally(() => setDetailLoading(false));
  };

  if (loading) return <Spinner animation="border" />;
  if (error)   return <Alert variant="danger">{error}</Alert>;

  // client-side pagination
  const pages = Math.ceil(list.length / perPage);
  const pageList = list.slice((page - 1) * perPage, page * perPage);

  return (
    <>
      <h2>{cfg.pluralLabel}</h2>
      <div className="mb-3 d-flex gap-2">
        <Button size="sm" onClick={() => setShowCreate(true)}>
          Create {cfg.singularLabel}
        </Button>
      </div>

      <ul className="list-group">
        {pageList.map(item => (
          <li
            key={item[idKey]}
            className="list-group-item list-group-item-action"
            style={{ cursor: 'pointer' }}
            onClick={() => openDetail(item[idKey])}
          >
            {item[idKey]} – {item[labelField]}
          </li>
        ))}
      </ul>

      <Pagination className="mt-3">
        <Pagination.Prev
          disabled={page === 1}
          onClick={() => setPage(p => p - 1)}
        />
        {[...Array(pages).keys()].map(i => (
          <Pagination.Item
            key={i + 1}
            active={i + 1 === page}
            onClick={() => setPage(i + 1)}
          >
            {i + 1}
          </Pagination.Item>
        ))}
        <Pagination.Next
          disabled={page === pages}
          onClick={() => setPage(p => p + 1)}
        />
      </Pagination>

      {/* Create Modal */}
      <Modal show={showCreate} onHide={() => setShowCreate(false)} size="xl">
        <Modal.Header closeButton>
          <Modal.Title>Create {cfg.singularLabel}</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form onSubmit={handleCreate}>
            <Row className="g-3">
              {newFields.map(f => (
                <Col key={f} xs={12}>
                  <Form.Group controlId={`form-${f}`}>
                    <Form.Label>
                      {f.replace(/_/g, ' ')}
                      {errors[f] && (
                        <span className="text-danger ms-2">{errors[f]}</span>
                      )}
                    </Form.Label>
                    {fkOpts[f] ? (
                      <Form.Select
                        value={formData[f]}
                        isInvalid={!!errors[f]}
                        onChange={e =>
                          setFormData(prev => ({ ...prev, [f]: e.target.value }))
                        }
                      >
                        <option value="">
                          {`Select ${f.replace(/_id/, '')}`}
                        </option>
                        {fkOpts[f].map(o => {
                          const endpoint = fkMap[f];
                          const labelKey = labelMap[endpoint];
                          return (
                            <option key={o[idKey]} value={o[idKey]}>
                              {`${o[idKey]} – ${o[labelKey]}`}
                            </option>
                          );
                        })}
                      </Form.Select>
                    ) : (
                      <Form.Control
                        type={schema[resourceKey][f] === 'int' ? 'number' : 'text'}
                        placeholder={getPlaceholder(f)}
                        value={formData[f]}
                        isInvalid={!!errors[f]}
                        onChange={e =>
                          setFormData(prev => ({ ...prev, [f]: e.target.value }))
                        }
                      />
                    )}
                  </Form.Group>
                </Col>
              ))}
            </Row>
            <div className="mt-3">
              <Button type="submit">Submit</Button>
            </div>
          </Form>
        </Modal.Body>
      </Modal>

      {/* Details Modal */}
      <Modal show={showDetail} onHide={() => setShowDetail(false)} size="lg">
        <Modal.Header closeButton>
          <Modal.Title>
            {cfg.singularLabel} Details (ID: {detailId})
          </Modal.Title>
        </Modal.Header>
        <Modal.Body>
          {detailLoading && <Spinner animation="border" />}
          {detailError && <Alert variant="danger">{detailError}</Alert>}
          {detail && (
            <dl className="row">
              {Object.entries(detail).map(([k, v]) => (
                <React.Fragment key={k}>
                  <dt className="col-sm-3">{k}</dt>
                  <dd className="col-sm-9">{String(v)}</dd>
                </React.Fragment>
              ))}
            </dl>
          )}
        </Modal.Body>
      </Modal>
    </>
  );
}
