const MAX = 100;

function main() {
    let divisi = new Array(MAX);
    let kuota = new Array(MAX);
    let terdaftar = new Array(MAX);

    divisi[0] = "Software Development";
    divisi[1] = "UI/UX Design";
    divisi[2] = "Technopreneur";
    divisi[3] = "Intelligence System";

    kuota[0] = 20;
    kuota[1] = 15;
    kuota[2] = 25;
    kuota[3] = 10;

    terdaftar[0] = 18;
    terdaftar[1] = 15;
    terdaftar[2] = 10;
    terdaftar[3] = 10;

    const nData = 4;

    console.log("Study Group yang Masih Tersedia:");
    for (let i = 0; i < nData; i++) {
        if (terdaftar[i] < kuota[i]) {
            console.log("-", divisi[i]);
        }
    }

    console.log("\nStudy Group yang Sudah Penuh:");
    for (let i = 0; i < nData; i++) {
        if (terdaftar[i] >= kuota[i]) {
            console.log("-", divisi[i]);
        }
    }
}

main();
