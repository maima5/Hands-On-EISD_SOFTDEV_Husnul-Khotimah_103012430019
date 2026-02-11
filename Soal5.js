class Peserta {
    constructor(divisi, jumlah) {
        this.Divisi = divisi;
        this.Jumlah = jumlah;
    }
}

function main() {
    const data = [
        new Peserta("Web Development Dasar", 18),
        new Peserta("Logic & Algorithm", 15),
        new Peserta("Frontend Fundamental", 10),
        new Peserta("Backend Introduction", 10)
    ];

    let total = 0;
    for (let i = 0; i < 4; i++) {
        total = total + data[i].Jumlah;
    }

    const rataRata = total / data.length;

    console.log("Total seluruh peserta:", total);
    console.log("Rata-rata peserta per Study Group:", rataRata);
}

main();
