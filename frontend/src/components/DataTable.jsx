import React from "react";

export default function DataTable({ mode, data }) {
  if (mode === "clubs") {
    return (
      <table border="1" width="100%">
        <thead>
          <tr>
            <th>ID</th>
            <th>Клуб</th>
            <th>Количество титулов</th>
            <th>Средний возраст</th>
            <th>Город</th>
          </tr>
        </thead>
        <tbody>
          {data.length === 0 ? (
            <tr><td colSpan="5" style={{ textAlign: "center" }}>Нет данных</td></tr>
          ) : (
            data.map((club) => (
              <tr key={club.id}>
                <td>{club.id}</td>
                <td>{club.name}</td>
                <td>{club.quantity_tituls}</td>
                <td>{club.average_age_players}</td>
                <td>{club.town?.name || club.town?.name_town}</td> {}
              </tr>
            ))
          )}
        </tbody>
      </table>
    );
  }

  if (mode === "towns") {
    return (
      <table border="1" width="100%">
        <thead>
          <tr>
            <th>ID</th>
            <th>Название города</th>
          </tr>
        </thead>
        <tbody>
          {data.length === 0 ? (
            <tr><td colSpan="2" style={{ textAlign: "center" }}>Нет данных</td></tr>
          ) : (
            data.map((town) => (
              <tr key={town.id}>
                <td>{town.id}</td>
                <td>{town.name || town.name_town}</td> {}
              </tr>
            ))
          )}
        </tbody>
      </table>
    );
  }

  return <p>Выберите режим</p>;
}