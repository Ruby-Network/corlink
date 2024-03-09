try {
     if (!localStorage["auth"] && new URL(document.all.rcheck.href).password) {
        window.location.reload();
        localStorage["auth"] = true;
    }
}
catch (e) {
    console.log(e);
}
