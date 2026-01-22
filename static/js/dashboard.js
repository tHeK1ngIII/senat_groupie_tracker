async function loadPlayer() {
    const tag = document.getElementById("playerTag").value.replace("#", "%23");
    const res = await fetch(`/api/player?tag=${tag}`);
    const data = await res.json();
    document.getElementById("playerResult").textContent = JSON.stringify(data, null, 2);
}
async function loadClan() {
    let tag = document.getElementById("clanTag").value.trim();

    // Vérification du format du tag (8-9 caractères valides)
    if (!tag.match(/^#?[0289PYLQGRJCUV]{8,9}$/)) {
        alert("Tag invalide !");
        return;
    }

    // Encodage du #
    tag = tag.replace("#", "%23");

    try {
        const res = await fetch(`/api/clan?tag=${tag}`);
        if (!res.ok) {
            throw new Error(`Erreur serveur: ${res.status}`);
        }

        const data = await res.json();

        // Affichage formaté
        document.getElementById("clanResult").textContent =
            JSON.stringify(data, null, 2);

    } catch (err) {
        console.error(err);
        document.getElementById("clanResult").textContent =
            "Impossible de charger les infos du clan.";
    }
}



async function loadCards() {
    const res = await fetch(`/api/cards`);
    const data = await res.json();
    console.log(data);
    document.getElementById("cardsResult").textContent = JSON.stringify(data, null, 2);
}

async function loadBattles() {
    const tag = document.getElementById("battleTag").value.replace("#", "%23");
    const res = await fetch(`/api/battles?tag=${tag}`);
    const data = await res.json();
    document.getElementById("battleResult").textContent = JSON.stringify(data, null, 2);
}

import React, { useEffect, useState } from "react";

function Dashboard() {
  const [clan, setClan] = useState(null);

  useEffect(() => {
    fetch("https://serene-encouragement.up.railway.app/api/clan?tag=%238CP0C2PL")
      .then(res => res.json())
      .then(data => setClan(data))
      .catch(err => console.error(err));
  }, []);

  return (
    <div>
      <h1>Clash Royale Dashboard</h1>
      {clan ? <pre>{JSON.stringify(clan, null, 2)}</pre> : <p>Chargement...</p>}
    </div>
  );
}

export default Dashboard;