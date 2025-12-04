import React from "react";

export default function ControlPanel({
  search,
  setSearch,
  searchType,
  setSearchType,
  onSearch,
  onShowClubs,
  onShowTowns,
  onCreateClub,
  onEditClub,
  onDeleteClub,
  onGetClub,
  onCreateTown,
  onEditTown,
  onDeleteTown,
  onGetTown
}) {
  return (
    <div style={panelStyle}>
      <select
        value={searchType}
        onChange={(e) => setSearchType(e.target.value)}
        style={selectStyle}
      >
        <option value="club">По клубу</option>
        <option value="town">По городу</option>
      </select>
      <input
        type="text"
        placeholder="Поиск..."
        value={search}
        onChange={(e) => setSearch(e.target.value)}
        style={searchStyle}
      />
      <button onClick={onSearch} style={btnSearch}>🔍</button>

      <div style={blockStyle}>
        <span style={label}>Клуб:</span>
        <button style={btnGreen} onClick={onCreateClub}>＋</button>
        <button style={btnGreen} onClick={onEditClub}>✎</button>
        <button style={btnGreen} onClick={onDeleteClub}>❌</button>
        <button style={btnGreen} onClick={onShowClubs}>👁</button>
        <button style={btnGreen} onClick={onGetClub}>📄</button>
      </div>

      <div style={blockStyle}>
        <span style={label}>Город:</span>
        <button style={btnOrange} onClick={onCreateTown}>＋</button>
        <button style={btnOrange} onClick={onEditTown}>✎</button>
        <button style={btnOrange} onClick={onDeleteTown}>❌</button>
        <button style={btnOrange} onClick={onShowTowns}>👁</button>
        <button style={btnOrange} onClick={onGetTown}>📄</button>
      </div>
    </div>
  );
}

const panelStyle = {
  display: "flex",
  alignItems: "center",
  gap: "12px",
  flexWrap: "wrap",
  marginBottom: "16px",
};

const selectStyle = {
  padding: "8px",
  borderRadius: "6px",
  border: "1px solid #ccc",
  fontSize: "14px",
};

const searchStyle = {
  padding: "8px",
  width: "220px",
  borderRadius: "6px",
  border: "1px solid #ccc",
  fontSize: "14px",
};

const btnSearch = {
  backgroundColor: "#4285F4",
  color: "white",
  border: "none",
  borderRadius: "6px",
  padding: "8px 12px",
  cursor: "pointer",
};

const label = {
  marginRight: "6px",
  fontWeight: "bold",
  color: "#333",
};

const blockStyle = {
  display: "flex",
  alignItems: "center",
  gap: "4px",
};

const baseButton = {
  width: "32px",
  height: "32px",
  border: "none",
  borderRadius: "6px",
  fontSize: "16px",
  color: "white",
  cursor: "pointer",
};

const btnGreen = { ...baseButton, backgroundColor: "#6ecb63" };
const btnOrange = { ...baseButton, backgroundColor: "#ffb74d" };