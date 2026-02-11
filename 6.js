function main() {
    const emails = [
        "Andi@gmail.com",
        "Budi@gmail.com",
        "Sari@gmail.com",
        "Andi@gmail.com",
        "Rina@gmail.com",
        "Budi@gmail.com"
    ];

    let adaDuplikat = false;
    let stop = false;
    console.log("Email duplikat:");

    for (let i = 0; i < 6; i++) {
        stop = false;
        for (let j = i + 1; j < 6 && stop === false; j++) {
            if (emails[i] === emails[j]) {
                console.log("-", emails[i]);
                adaDuplikat = true;
                stop = true;
            }
        }
    }

    if (!adaDuplikat) {
        console.log("Tidak ada data duplikat");
    }
}

main();