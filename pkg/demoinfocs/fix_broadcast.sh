#!/bin/bash
offset=0

truncate -s 0 new.dem

while read line; do

    end=$(echo $line | cut -f1 -d,)
    len=$(($end - $offset))
#    tail -c +$(($offset + 1)) fake.dem | head -c $len >> new.dem
#    cat command_info.bin >> new.dem
    dd if=fake.dem iflag=count_bytes,skip_bytes skip=$offset count=$len status=none >> new.dem
    dd if=command_info.bin status=none >> new.dem

    offset=$end

done < dc_offsets.txt

dd if=fake.dem iflag=skip_bytes skip=$offset status=none >> new.dem

fake=$(wc -c fake.dem | cut -f1 -d' ')
echo "fake.dem: $fake"

echo "new.dem: $(wc -c new.dem | cut -f1 -d' ')"

dcs=$(wc -l dc_offsets.txt | cut -f1 -d' ')
echo "supposed to be $(($fake + ($dcs * 160)))"
