// Allows custom dropdowns 
function enableDropdown() {
    document.querySelectorAll(".dropdown").forEach(dropdown => {
        const selected = dropdown.querySelector(".dropdown__selected");
        updatedSelectedOption(dropdown, selected)
        toggleDropdown(dropdown, selected)
  });
}

// Updates selected value and selected text
function updatedSelectedOption(dropdown, selected) {
    const selectedText = selected.querySelector("p");
    const input = dropdown.querySelector("input");

    dropdown.querySelectorAll(".dropdown__options li").forEach(li => {
        li.addEventListener("click", e => {
            e.preventDefault();
            selectedText.textContent = li.textContent;
            input.value = li.dataset.value;
            updatedBackground(li, dropdown)
        });
    });
}

// Allows option to be hidden and shown and signifier to change
function toggleDropdown(dropdown, selected) {
    const selectedSignifier = selected.querySelector("span");
    selected.addEventListener("click", () => { 
        if (selectedSignifier.classList.contains("dropdown__signifier--expanded")) {
            selectedSignifier.classList.remove("dropdown__signifier--expanded");
        } else {
            selectedSignifier.classList.add("dropdown__signifier--expanded");
        }
        dropdown.querySelectorAll("li").forEach(li => {
        li.classList.toggle("dropdown__options--hidden");
        });
    }); 
}

// Highlights selected option background
function updatedBackground(clickedLI, dropdown) {
    if (clickedLI.classList.add("dropdown__option--selected")) {
        return;
    }
    dropdown.querySelectorAll(".dropdown__options li").forEach(li => {
        if (li.classList.contains("dropdown__option--selected")) {
            li.classList.remove("dropdown__option--selected");
        }     
    });
    clickedLI.classList.add("dropdown__option--selected");
}

enableDropdown()