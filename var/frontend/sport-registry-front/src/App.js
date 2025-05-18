
import React, { useState } from 'react';
import { Container, Row, Col } from 'react-bootstrap';
import Sidebar from './components/Sidebar';
import ResourceView from './components/ResourceView';

const resources = [
  { key: 'sports',             pluralLabel: 'Sports',             singularLabel: 'Sport' },
  { key: 'countries',          pluralLabel: 'Countries',          singularLabel: 'Country' },
  { key: 'currencies',         pluralLabel: 'Currencies',         singularLabel: 'Currency' },
  { key: 'locations',          pluralLabel: 'Locations',          singularLabel: 'Location' },
  { key: 'roles',              pluralLabel: 'Roles',              singularLabel: 'Role' },
  { key: 'competition-levels', pluralLabel: 'Competition Levels', singularLabel: 'Competition Level' },
  {
    key: 'persons', pluralLabel: 'People', singularLabel: 'Person', children: [
      { key: 'persons/sport', pluralLabel: 'Sports',       singularLabel: 'Sport' },
      { key: 'persons/team',  pluralLabel: 'Team Persons', singularLabel: 'Team Person' },
    ]
  },
  { key: 'teams',             pluralLabel: 'Teams',             singularLabel: 'Team' },
  { key: 'competitions',      pluralLabel: 'Competitions',      singularLabel: 'Competition' },
  { key: 'stages',            pluralLabel: 'Stages',            singularLabel: 'Stage' },
  { key: 'matches',           pluralLabel: 'Matches',           singularLabel: 'Match' },
  { key: 'match-participants',pluralLabel: 'Match Participants',singularLabel: 'Match Participant' },
  { key: 'prizes',            pluralLabel: 'Prizes',            singularLabel: 'Prize' },
  { key: 'competition-teams', pluralLabel: 'Competition Teams', singularLabel: 'Competition Team' },
  { key: 'team-achievements', pluralLabel: 'Team Achievements', singularLabel: 'Team Achievement' },
];

export default function App() {
  const [current, setCurrent] = useState(resources[0].key);
  return (
    <Container fluid className="p-0">
      <Row>
        <Col xs={12} md={3} lg={2}
             className="bg-light border-end vh-100 position-fixed"
             style={{ overflowY: 'auto' }}>
          <Sidebar resources={resources} current={current} onSelect={setCurrent} />
        </Col>
        <Col xs={12} md={{ span: 9, offset: 3 }} lg={{ span: 10, offset: 2 }}
             className="p-4"
             style={{ maxHeight: '100vh', overflowY: 'auto' }}>
          <ResourceView resourceKey={current} resources={resources} />
        </Col>
      </Row>
    </Container>
  );
}