async function loadPlayer() {
    const tag = document.getElementById("playerTag").value.replace("#", "%23");
    const res = await fetch(`/api/player?tag=${tag}`);
    const data = await res.json();
    document.getElementById("playerResult").textContent = JSON.stringify(data, null, 2);
}

async function loadClan() {
    const tag = document.getElementById("clanTag").value.replace("#", "%23");
    const res = await fetch(`/api/clan?tag=${tag}`);
    const data = await res.json();
    document.getElementById("clanResult").textContent = JSON.stringify(data, null, 2);
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