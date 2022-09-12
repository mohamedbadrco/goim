
for i in  *.gif;do ffmpeg -i "$i" "i/frame$i 0%d.png"  ; done

for i in  *.gif;do for j in $(seq 1 30); do rm "i/frame$i 0$j.png"; done; done


ffmpeg -framerate 14 -pattern_type glob -i 'i/*.png' -c:a copy -shortest -c:v libx264 -pix_fmt yuv420p out1.mp4

ffmpeg -f concat -i videos.txt -c copy o.mp4

rm out1.mp4

ffmpeg -i o.mp4 -i "sub/$(ls sub | shuf -n 1)" -c:v copy -c:a aac output.mp4

rm o.mp4

cd i

for i in  *.png;do rm "$i";done

cd ..



for i in  *.gif;do rm "$i";done