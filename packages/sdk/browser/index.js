try {
     if (!localStorage["auth"] && new URL(document.all.rcheck.href).password) {
        window.location.reload();
        localStorage["auth"] = true;
    }
}
catch {
    // do nothing
}
