filename = ARGV[0]
data = File.readlines(filename)
uniques = data.map(&:split).map(&:first).uniq.count

puts "there were #{uniques} visitors to your blog"
