
document.addEventListener("DOMContentLoaded", function() {
    let itemLinks = document.getElementsByClassName("item-link");
    let url = "localhost:8080"
    if (window.location.href.includes("jamfesteq.com")) {
        url = "play.jamfesteq.com"
    }

    for (let i = 0; i < itemLinks.length; i++) {
        let itemLink = itemLinks[i];
        let description = document.createElement("div");
        description.setAttribute("class", "item-description");
        description.setAttribute("style", "visibility: hidden;");
        description.innerHTML = "<img src='http://"+url+"/item/preview.png?id=" + itemLink.getAttribute("item-id") + "'>"
        itemLink.appendChild(description);

        itemLink.addEventListener("mouseover", (e) => {
            console.log("hovered in");
            let description = itemLink.getElementsByClassName("item-description")[0]
            if (description == null) {
                return;
            }
            description.setAttribute("style", "visibility: visible;");
            e.preventDefault();
        });

        itemLink.addEventListener("mouseout", (e) => {
            console.log("hovered out");
            let description = itemLink.getElementsByClassName("item-description")[0]
            if (description == null) {
                return;
            }
            description.setAttribute("style", "visibility: hidden;");
            e.preventDefault();
        });

        itemLink.addEventListener("click", (e) => {
            console.log("clicked");
            let description = itemLink.getElementsByClassName("item-description")[0]
            
            e.preventDefault();
        });
    }

    let spellLinks = document.getElementsByClassName("spell-link");

    for (let i = 0; i < spellLinks.length; i++) {
        let spellLink = spellLinks[i];
        let description = document.createElement("div");
        description.setAttribute("class", "spell-description");
        description.setAttribute("style", "visibility: hidden;");
        description.innerHTML = "<img src='http://"+url+"/spell/preview.png?id=" + spellLink.getAttribute("spell-id") + "'>"
        spellLink.appendChild(description);

        spellLink.addEventListener("mouseover", (e) => {
            let description = spellLink.getElementsByClassName("spell-description")[0]
            if (description == null) {
                return;
            }
            description.setAttribute("style", "visibility: visible;");
            e.preventDefault();
        });

        spellLink.addEventListener("mouseout", (e) => {
            let description = spellLink.getElementsByClassName("spell-description")[0]
            if (description == null) {
                return;
            }
            description.setAttribute("style", "visibility: hidden;");
            e.preventDefault();
        });

        // spellLink.addEventListener("click", (e) => {
        //     console.log("clicked");
        //     let description = spellLink.getElementsByClassName("spell-description")[0]
            
        //     e.preventDefault();
        // });
    }
});