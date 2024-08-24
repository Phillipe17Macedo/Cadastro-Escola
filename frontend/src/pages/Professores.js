import React, { useState, useEffect } from 'react';
import axios from 'axios';

function Professores() {
  const [professores, setProfessores] = useState([]);
  const [novoProfessor, setNovoProfessor] = useState({ nome: '', email: '', cpf: '' });

  useEffect(() => {
    axios.get('http://localhost:8080/professores')
      .then(response => setProfessores(response.data))
      .catch(error => console.log(error));
  }, []);

  const handleChange = (e) => {
    setNovoProfessor({ ...novoProfessor, [e.target.name]: e.target.value });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    console.log('Submitting the form');
    axios.post('http://localhost:8080/professores', novoProfessor)
      .then(response => {
        setProfessores([...professores, response.data]);
        setNovoProfessor({ nome: '', email: '', cpf: '' });
      })
      .catch(error => console.log(error));
  };

  return (
    <div className="container">
      <h2 className="my-4">Professores</h2>

      <h3 className="mb-3">Cadastrar Novo Professor</h3>
      <form onSubmit={handleSubmit} className="mb-4" noValidate>
        <div className="mb-3">
          <label htmlFor="nome" className="form-label">Nome</label>
          <input
            type="text"
            className="form-control"
            name="nome"
            id="nome"
            placeholder="Nome"
            value={novoProfessor.nome}
            onChange={handleChange}
            required
          />
        </div>
        <div className="mb-3">
          <label htmlFor="email" className="form-label">Email</label>
          <input
            type="email"
            className="form-control"
            name="email"
            id="email"
            placeholder="Email"
            value={novoProfessor.email}
            onChange={handleChange}
            required
          />
        </div>
        <div className="mb-3">
          <label htmlFor="cpf" className="form-label">CPF</label>
          <input
            type="text"
            className="form-control"
            name="cpf"
            id="cpf"
            placeholder="CPF"
            value={novoProfessor.cpf}
            onChange={handleChange}
            required
          />
        </div>
        <button type="submit" className="btn btn-primary">Cadastrar</button>
      </form>

      <h3 className="mb-3">Lista de Professores</h3>
      <ul className="list-group">
        {professores.map(professor => (
          <li key={professor.ID} className="list-group-item">
            {professor.Nome} - {professor.Email}
          </li>
        ))}
      </ul>
    </div>
  );
}

export default Professores;