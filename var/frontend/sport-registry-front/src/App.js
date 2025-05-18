import React, { useState } from 'react';
import { Container, Row, Col } from 'react-bootstrap';
import Sidebar from './components/Sidebar';
import Dashboard from './components/Dashboard';
import ResourceView from './components/ResourceView';

const resources = [
  { key: 'dashboard',         pluralLabel: 'Dashboard',           singularLabel: '' },
  { key: 'sports',            pluralLabel: 'Sports',              singularLabel: 'Sport' },
  { key: 'countries',         pluralLabel: 'Countries',           singularLabel: 'Country' },
  { key: 'currencies',        pluralLabel: 'Currencies',          singularLabel: 'Currency' },
  { key: 'locations',         pluralLabel: 'Locations',           singularLabel: 'Location' },
  { key: 'roles',             pluralLabel: 'Roles',               singularLabel: 'Role' },
  {
    key: 'persons',
    pluralLabel: 'People',
    singularLabel: 'Person',
    children: [
      { key: 'persons/sport', pluralLabel: 'Sports',       singularLabel: 'Sport' },
      { key: 'persons/team',  pluralLabel: 'Team Persons', singularLabel: 'Team Person' },
    ]
  },
  {
    key: 'teams',
    pluralLabel: 'Teams',
    singularLabel: 'Team',
    children: [
      { key: 'competition-teams',  pluralLabel: 'Competition Teams',  singularLabel: 'Competition Team' },
      { key: 'team-achievements',  pluralLabel: 'Team Achievements',  singularLabel: 'Team Achievement' },
      { key: 'match-participants', pluralLabel: 'Match Participants', singularLabel: 'Match Participant' },
    ]
  },
  {
    key: 'competitions',
    pluralLabel: 'Competitions',
    singularLabel: 'Competition',
    children: [
      { key: 'competition-levels', pluralLabel: 'Competition Levels', singularLabel: 'Competition Level' },
      { key: 'stages', pluralLabel: 'Stages', singularLabel: 'Stage' },
      { key: 'matches', pluralLabel: 'Matches', singularLabel: 'Match' }
    ]
  },
  { key: 'prizes',            pluralLabel: 'Prizes',              singularLabel: 'Prize' },
];

export default function App() {
  // По умолчанию показываем Dashboard
  const [current, setCurrent] = useState('dashboard');

  return (
    <Container fluid className="p-0">
      <Row>
        {/* Сайдбар */}
        <Col
          xs={12} md={3} lg={2}
          className="bg-light border-end vh-100 position-fixed"
          style={{ overflowY: 'auto' }}
        >
          <Sidebar
            resources={resources}
            current={current}
            onSelect={setCurrent}
          />
        </Col>

        {/* Основной контент */}
        <Col
          xs={12}
          md={{ span: 9, offset: 3 }}
          lg={{ span: 10, offset: 2 }}
          className="p-4"
          style={{ maxHeight: '100vh', overflowY: 'auto' }}
        >
          {current === 'dashboard' ? (
            <Dashboard resources={resources.filter(r => r.key !== 'dashboard')} />
          ) : (
            <ResourceView resourceKey={current} resources={resources} />
          )}
        </Col>
      </Row>
    </Container>
  );
}
