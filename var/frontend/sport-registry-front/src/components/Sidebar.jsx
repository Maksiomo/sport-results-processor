import React, { useState } from 'react';
import { Nav, Form } from 'react-bootstrap';

export default function Sidebar({ resources, current, onSelect }) {
  const [filter, setFilter] = useState('');

  // Фильтруем и включаем вложенные, если родитель проходит
  const filtered = resources
    .map(res => {
      // проверим сам ресурс
      const matchParent = res.pluralLabel.toLowerCase().includes(filter.toLowerCase());
      // если есть дети — тоже фильтруем их
      const children = res.children
        ? res.children.filter(c =>
            c.pluralLabel.toLowerCase().includes(filter.toLowerCase())
          )
        : null;
      if (matchParent || (children && children.length)) {
        return { ...res, children };
      }
      return null;
    })
    .filter(Boolean);

  return (
    <div className="p-3">
      <Form.Control
        type="text"
        placeholder="Filter..."
        className="mb-3"
        value={filter}
        onChange={e => setFilter(e.target.value)}
      />
      <Nav variant="pills" className="flex-column" activeKey={current}>
        {filtered.map(res => (
          <React.Fragment key={res.key}>
            <Nav.Item>
              <Nav.Link
                eventKey={res.key}
                onClick={() => onSelect(res.key)}
                className="fw-bold"
              >
                {res.pluralLabel}
              </Nav.Link>
            </Nav.Item>
            {res.children && res.children.map(child => (
              <Nav.Item key={child.key}>
                <Nav.Link
                  eventKey={child.key}
                  onClick={() => onSelect(child.key)}
                  className="ps-4"
                >
                  {child.pluralLabel}
                </Nav.Link>
              </Nav.Item>
            ))}
          </React.Fragment>
        ))}
        {filtered.length === 0 && (
          <div className="text-muted fst-italic">No matches</div>
        )}
      </Nav>
    </div>
  );
}
