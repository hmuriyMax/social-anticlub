create index
    if not exists user_info_names
    on user_info(first_name text_pattern_ops, second_name text_pattern_ops);