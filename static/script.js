let folderStack = [];
let currentFolder = "root";

async function loadFiles(folderId = "root") {
    currentFolder = folderId;

    const res = await fetch(`/drive/files?folderId=${folderId}`);
    const files = await res.json();

    const container = document.getElementById("files");
    container.innerHTML = "";

    files.forEach(file => {
        const div = document.createElement("div");

        if (file.mimeType === "application/vnd.google-apps.folder") {
            div.innerText = "📁 " + file.name;
            div.className = "folder";
            div.onclick = () => {
                folderStack.push(currentFolder);
                loadFiles(file.id);
            };
        } else {
            div.innerText = "📄 " + file.name;
            div.className = "file";
        }

        container.appendChild(div);
    });
}

function goBack() {
    if (folderStack.length > 0) {
        const prev = folderStack.pop();
        loadFiles(prev);
    }
}

window.onload = () => {
    loadFiles();
};