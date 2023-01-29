let SelectedCard

for (let i = 1; i <= 40; i++) {
    document.getElementById(String(i)).addEventListener("click", moveCard)
}

function moveCard(e) {
    if (SelectedCard != null) {
        e.target.appendChild(SelectedCard)
        SelectedCard = null
    }

    e.stopPropagation()
}

document.getElementById("shkilet").addEventListener("click", selectCard)
document.getElementById("soldier").addEventListener("click", selectCard)

function selectCard(e) {
    SelectedCard = e.target
    e.stopPropagation()
}

