
document.addEventListener("DOMContentLoaded", function() {
    let url = "localhost:8080"
    if (window.location.href.includes("jamfesteq.com")) {
        url = "play.jamfesteq.com"
    }

    let links = ['item', 'spell', 'npc'];


    for (let i = 0; i < links.length; i++) {
        let linkName = links[i];
        console.log("searching for", linkName);
        let elementLinks = document.getElementsByClassName(linkName+"-link");

        for (let j = 0; j < elementLinks.length; j++) {
            let elementLink = elementLinks[j];
            if (elementLink == null) {
                continue;
            }
            console.log("Found element", elementLink);
            let description = document.createElement("div");
            description.setAttribute("class", linkName+"-description");
            description.setAttribute("style", "visibility: hidden;");
            elementLink.appendChild(description);

            
            
            elementLink.addEventListener("mouseover", (e) => {
                console.log("hovered in");
                let description = elementLink.getElementsByClassName(linkName+"-description")[0]
                if (description == null) {
                    return;
                }

                // check if description has img element
                let img = description.getElementsByTagName("img")[0];
                if (img == null) {
                    description.innerHTML = "<img src='http://"+url+"/"+linkName+"/preview.png?id=" + elementLink.getAttribute(linkName+"-id") + "'>"
                    elementLink.appendChild(description);
                }
                if (description.getAttribute("clicked") == "true") {
                    return;
                }


                description.setAttribute("style", "visibility: visible;");
                e.preventDefault();
            });

            elementLink.addEventListener("mouseout", (e) => {
                console.log("hovered out");
                let description = elementLink.getElementsByClassName(linkName+"-description")[0]
                if (description == null) {
                    return;
                }
                if (description.getAttribute("clicked") == "true") {
                    return;
                }
                description.setAttribute("style", "visibility: hidden;");
                e.preventDefault();
            });

            elementLink.addEventListener("click", (e) => {
                let description = elementLink.getElementsByClassName(linkName+"-description")[0]
                if (description == null) {
                    return;
                }
                e.preventDefault();
                let isVisible = description.style.visibility == "visible";
                console.log("clicked", description.style.visibility);
                if (description.hasAttribute("clicked") && isVisible) {
                    description.removeAttribute("clicked")
                    description.style.visibility = "hidden";
                    console.log("hiding");
                    return;
                }
                if (!isVisible) {
                    description.style.visibility = "visible";
                }
                description.setAttribute("clicked", "true")
            });
        }
    }
});