import re

def validasi_data(nama, email, pilihan, jumlah):
    valid = True

    if not nama or len(nama.strip()) < 3:
        valid = False
        
    #mengecek email dengan metode REGEX
    pola_email = r'^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$'
    if not email or not re.match(pola_email, email):
        valid = False

    if not pilihan or pilihan.strip() == "":
        valid = False

    if jumlah < 1:
        valid = False

    return valid


nama = input("Masukkan nama peserta: ")
email = input("Masukkan email: ")
pilihan = input("Masukkan pilihan study group: ")

try:
    jumlah = int(input("Masukkan jumlah peserta terdaftar: "))
except ValueError:
    jumlah = 0

if validasi_data(nama, email, pilihan, jumlah):
    print("Data valid")
else:
    print("Data tidak valid")
