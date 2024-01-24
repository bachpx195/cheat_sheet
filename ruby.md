### Group and count

def group_and_count arr, sort_by_value=true
  h = arr.inject(Hash.new(0)) { |h, e| h[e] += 1 ; h }
  if sort_by_value
    return h.sort_by {|_key, value| value}.to_h
  else
    return h.sort.to_h
  end
end
