import React, { useState, useEffect } from 'react'
import {
  Table, Form, Button, Row, Col, Alert,
  Modal, Spinner
} from 'react-bootstrap'
import api from '../api'

export default function ResourceView({ resourceKey, resources }) {
  // 1. Метаданные (labels, children)
  const findMeta = (key, list) => {
    for (const r of list) {
      if (r.key === key) return r
      if (r.children) {
        const c = findMeta(key, r.children)
        if (c) return c
      }
    }
    return null
  }
  const cfgMeta = findMeta(resourceKey, resources)
  if (!cfgMeta) return <Alert variant="warning">Unknown resource: {resourceKey}</Alert>

  // 2. Типы полей
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
  }

  // 3. Поля для таблицы и формы
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
  }
  const { fields, newFields } = fieldMeta[resourceKey]

  // 4. Настройки внешних ключей
  const fkMap = {
    country_id: 'countries',
    sport_id: 'sports',
    person_id: 'persons',
    team_id: 'teams',
    role_id: 'roles',
    competition_id: 'competitions',
    stage_id: 'stages',
    match_id: 'matches',
    prize_id: 'prizes',
    location_id: 'locations',
    currency_code: 'currencies',
  }
  const labelMap = {
    countries: 'name',
    sports: 'name',
    persons: 'name',
    teams: 'name',
    roles: 'name',
    competitions: 'name',
    stages: 'name',
    matches: 'match_time',
    prizes: 'place_bracket',
    locations: 'full_address',
    currencies: 'name',
  }
  const primaryKeyMap = { currencies: 'code', default: 'id' }
  const idKey = primaryKeyMap[resourceKey] || primaryKeyMap.default

  // 5. Состояния
  const [items, setItems]             = useState([])
  const [loading, setLoading]         = useState(true)
  const [error, setError]             = useState(null)

  // Create modal
  const [showCreate, setShowCreate]   = useState(false)
  const [formData, setFormData]       = useState(newFields.reduce((a,f)=>(a[f]='',a),{}))
  const [foreignOpts, setForeignOpts] = useState({})

  // Get modal
  const [showGet, setShowGet]         = useState(false)
  const [getId, setGetId]             = useState('')
  const [getResult, setGetResult]     = useState(null)
  const [getError, setGetError]       = useState(null)
  const [getLoading, setGetLoading]   = useState(false)

  // List modal
  const [showList, setShowList]       = useState(false)

  // 6. Генерация placeholder
  const getPlaceholder = f => {
    const t = schema[resourceKey][f]
    if (t==='int')     return 'e.g. 123'
    if (t==='boolean') return 'true or false'
    if (t==='json')    return '{"key":"value"}'
    return `Enter ${f.replace(/_/g,' ')}`
  }

  // 7. Загрузка данных
  useEffect(() => {
    setLoading(true)
    api.get(`/${resourceKey}`)
      .then(r => setItems(r.data.data ?? r.data))
      .catch(e => setError(e.message))
      .finally(() => setLoading(false))

    newFields.forEach(f => {
      const ep = fkMap[f]; if (!ep) return
      api.get(`/${ep}`)
        .then(r => setForeignOpts(prev => ({ ...prev, [f]: r.data.data ?? r.data })))
        .catch(console.error)
    })

    setFormData(newFields.reduce((a,f)=>(a[f]='',a),{}))
  }, [resourceKey])

  const convert = (k,v) => {
    const t = schema[resourceKey][k]
    if (t==='int')     return parseInt(v,10)
    if (t==='boolean') return v==='true'
    if (t==='json')    {
      try { return JSON.parse(v) } catch { return {} }
    }
    return v
  }

  const handleCreate = e => {
    e.preventDefault()
    const payload = {}
    newFields.forEach(f => payload[f] = convert(f, formData[f]))
    api.post(`/${resourceKey}`, payload)
      .then(r => {
        const obj = r.data.data ?? r.data
        setItems(prev => [...prev, obj])
        setShowCreate(false)
      })
      .catch(err => alert('Create error: '+err.message))
  }

  const handleGet = e => {
    e.preventDefault()
    setGetLoading(true)
    setGetError(null)
    setGetResult(null)
    api.get(`/${resourceKey}/${getId}`)
      .then(r => setGetResult(r.data.data ?? r.data))
      .catch(err => setGetError(err.message))
      .finally(() => setGetLoading(false))
  }

  if (loading) return <Spinner animation="border" />
  if (error)   return <Alert variant="danger">{error}</Alert>

  return (
    <>
      <h2>{cfgMeta.pluralLabel}</h2>
      <div className="mb-3 d-flex gap-2">
        <Button size="sm" onClick={()=>setShowList(true)}>
          List {cfgMeta.pluralLabel}
        </Button>
        <Button size="sm" onClick={()=>setShowCreate(true)}>
          Create {cfgMeta.singularLabel}
        </Button>
        <Button size="sm" onClick={()=>setShowGet(true)}>
          Get {cfgMeta.singularLabel}
        </Button>
      </div>

      {/* List Modal */}
      <Modal show={showList} onHide={()=>setShowList(false)} size="xl">
        <Modal.Header closeButton>
          <Modal.Title>List {cfgMeta.pluralLabel}</Modal.Title>
        </Modal.Header>
        <Modal.Body style={{ overflowX: 'auto' }}>
          <Table size="sm" striped bordered hover>
            <thead>
              <tr>{fields.map(f=><th key={f}>{f.replace(/_/g,' ')}</th>)}</tr>
            </thead>
            <tbody>
              {items.map(item=>(
                <tr key={item[idKey] ?? JSON.stringify(item)}>
                  {fields.map(f=><td key={f}>{String(item[f]??'')}</td>)}
                </tr>
              ))}
            </tbody>
          </Table>
        </Modal.Body>
      </Modal>

      {/* Create Modal */}
      <Modal show={showCreate} onHide={()=>setShowCreate(false)} size="lg">
        <Modal.Header closeButton>
          <Modal.Title>Create {cfgMeta.singularLabel}</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form onSubmit={handleCreate}>
            <Row className="g-3">
              {newFields.map(f=>{
                const opts = foreignOpts[f]
                return (
                  <Col key={f} xs={12}>
                    <Form.Group controlId={`form-${f}`}>
                      <Form.Label>{f.replace(/_/g,' ')}</Form.Label>
                      {opts ? (
                        <Form.Select
                          value={formData[f]}
                          onChange={e=>setFormData(p=>({...p,[f]:e.target.value}))}
                        >
                          <option value="">
                            {`Select ${cfgMeta.singularLabel}`}
                          </option>
                          {opts.map(o=>(
                            <option key={o[idKey]} value={o[idKey]}>
                              {`${o[idKey]} – ${o[labelMap[fkMap[f]]]}`}
                            </option>
                          ))}
                        </Form.Select>
                      ) : (
                        <Form.Control
                          type="text"
                          placeholder={getPlaceholder(f)}
                          value={formData[f]}
                          onChange={e=>setFormData(p=>({...p,[f]:e.target.value}))}
                        />
                      )}
                    </Form.Group>
                  </Col>
                )
              })}
            </Row>
            <div className="mt-3">
              <Button type="submit">Submit</Button>
            </div>
          </Form>
        </Modal.Body>
      </Modal>

      {/* Get Modal */}
      <Modal
        show={showGet}
        onHide={() => {
          setShowGet(false);
          setGetId('');
          setGetResult(null);
          setGetError(null);
        }}
        size="lg"
      >
        <Modal.Header closeButton>
          <Modal.Title>Get {cfgMeta.singularLabel}</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form onSubmit={handleGet}>
            <Form.Group controlId="form-get-id" className="mb-3">
              <Form.Label>{idKey.toUpperCase()}</Form.Label>
              <Form.Select
                value={getId}
                onChange={e => setGetId(e.target.value)}
              >
                <option value="">
                  {`Select ${cfgMeta.singularLabel}`}
                </option>
                {items.map(o => (
                  <option key={o[idKey]} value={o[idKey]}>
                    {`${o[idKey]} – ${o[labelMap[resourceKey]] ?? o[idKey]}`}
                  </option>
                ))}
              </Form.Select>
            </Form.Group>
            <Button type="submit">Fetch</Button>
          </Form>
          <hr />
          {getLoading && <Spinner animation="border" />}
          {getError && <Alert variant="danger">{getError}</Alert>}
          {getResult && (
            <pre className="bg-light p-2">
              {JSON.stringify(getResult, null, 2)}
            </pre>
          )}
        </Modal.Body>
      </Modal>
    </>
  )
}