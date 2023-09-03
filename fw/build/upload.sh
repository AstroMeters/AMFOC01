echo ${1:"a"}
python3 ../upload.py --serial ${1:/dev/ttyACM0} --family RP2040 --deploy dispenser.uf2
