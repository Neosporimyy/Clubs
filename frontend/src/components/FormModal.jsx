import React, { useState } from "react";

export default function FormModal({ mode, action, onSubmit, onCancel }) {
  const [formData, setFormData] = useState({});

  const handleChange = (e) => {
    const value = e.target.type === 'number' ? parseFloat(e.target.value) : e.target.value;
    setFormData({ ...formData, [e.target.name]: value });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    onSubmit(formData);
  };

  const titleMap = { create: "Создать", edit: "Редактировать", delete: "Удалить", get: "Получить по ID" };

  return (
    <div style={{ border: "1px solid #ccc", borderRadius: "8px", padding: "15px", marginBottom: "20px", backgroundColor: "#f9f9f9" }}>
      <h3>{titleMap[action]} {mode === "clubs" ? "клуб" : "город"}</h3>
      <form onSubmit={handleSubmit}>
        {(action === "edit" || action === "delete" || action === "get") && (
          <div style={{ marginBottom: "10px" }}>
            <input name="id" type="number" placeholder="ID" onChange={handleChange} style={inputStyle} required />
          </div>
        )}
        {action !== "delete" && action !== "get" && (
          <>
            {mode === "clubs" && (
              <>
                <div style={{ marginBottom: "10px" }}>
                  <input name="name" placeholder="Название клуба" onChange={handleChange} style={inputStyle} required />
                </div>
                <div style={{ marginBottom: "10px" }}>
                  <input name="quantity_tituls" type="number" placeholder="Количество титулов" onChange={handleChange} style={inputStyle} required />
                </div>
                <div style={{ marginBottom: "10px" }}>
                  <input name="average_age_players" type="number" step="0.1" placeholder="Средний возраст игроков" onChange={handleChange} style={inputStyle} required />
                </div>
                <div style={{ marginBottom: "10px" }}>
                  <input name="town_id" type="number" placeholder="ID города" onChange={handleChange} style={inputStyle} required />
                </div>
              </>
            )}
            {mode === "towns" && (
              <div style={{ marginBottom: "10px" }}>
                <input name="name" placeholder="Название города" onChange={handleChange} style={inputStyle} required />
              </div>
            )}
          </>
        )}
        <button type="submit" style={btnPrimary}>{titleMap[action]}</button>
        <button type="button" onClick={onCancel} style={{ ...btnPrimary, backgroundColor: "gray", marginLeft: "10px" }}>Отмена</button>
      </form>
    </div>
  );
}

const inputStyle = { padding: "8px", width: "300px", borderRadius: "5px", border: "1px solid #ccc" };
const btnPrimary = { padding: "8px 16px", borderRadius: "5px", border: "none", backgroundColor: "#007bff", color: "white", cursor: "pointer" };