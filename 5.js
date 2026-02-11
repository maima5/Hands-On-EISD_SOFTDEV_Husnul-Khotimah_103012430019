function main() {
    const peserta = [18, 15, 10, 10];

    let total = 0;
    for (let i = 0; i < 4; i++) {
        total = total + peserta[i];
    }

    const rataRata = total / peserta.length;

    console.log("Total seluruh peserta:", total);
    console.log("Rata-rata peserta per Study Group:", rataRata);
}

main();