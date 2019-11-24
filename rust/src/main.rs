#[macro_use]
extern crate serde_derive;
#[macro_use]
extern crate serde_json;

use std::env::args;
use std::fs;
use std::io;
use std::io::Write;
use std::path::Path;

mod problem;

const LEETCODE_NUMS: u32 = 1106;

fn main() {
    let mut solved_pros = get_solved_pros();
    let mut id_arg = String::new();
    let mut id: u32 = 0;
    //    io::stdin().read_line(&mut id_arg).expect("failed to read line");
    let args: Vec<String> = args().collect();
    if args.len() != 2 {
        panic!("args need be random or 1 - {}", LEETCODE_NUMS)
    }
    let id_arg = &args[1];
    let random = String::from("random");
    match id_arg {
        random => id = generate_random_id(&solved_pros),
        _ => {
            id = id_arg
                .parse::<u32>()
                .unwrap_or_else(|_| panic!("not a number {}", id_arg));
            if solved_pros.contains(&id) {
                println!("solved {}", id);
                return;
            }
        }
    }
    let p =
        problem::get_problem(id).unwrap_or_else(|| panic!("Error: failed to get problem #{} ", id));
    let code = p.code_definition.iter().find(|&d| d.value == "rust");
    if code.is_none() {
        panic!("problem {} has no rust version", &id);
    }
    let code = code.unwrap();
    let file_name = format!("n{:04}_{}", p.question_id, p.title_slug.replace("-", "_"));
    let file_path = Path::new("./src").join(&file_name);
    let template = fs::read_to_string("./problem.tmpl").unwrap();
    let source = template
        .replace("__PROBLEM_TITLE__", &p.title)
        .replace("__PROBLEM_DESC__", &build_desc(&p.content))
        .replace("__PROBLEM_DEFAULT_CODE__", &code.default_code)
        .replace("__PROBLEM_ID__", &format!("{}", p.question_id))
        .replace("__EXTRA_USE__", &parse_extra_use(&code.default_code));
    let mut file = fs::OpenOptions::new()
        .write(true)
        .create(true)
        .truncate(true)
        .open(&file_path)
        .unwrap();

    file.write_all(source.as_bytes()).unwrap();
    drop(file);
    let mut lib_file = fs::OpenOptions::new()
        .write(true)
        .append(true)
        .open("./src/lib.rs")
        .unwrap();
    writeln!(lib_file, "mod {};", file_name);
}

fn get_solved_pros() -> Vec<u32> {
    let paths = fs::read_dir("./src").unwrap();
    let mut res = Vec::new();
    for path in paths {
        let path = path.unwrap().path();
        let s = path.to_str().unwrap();
        if !s.starts_with('n') {
            continue;
        }
        let id = &s[7..11];
        let id = id.parse::<u32>().unwrap();
        res.push(id);
    }
    res
}

fn parse_extra_use(code: &str) -> String {
    let mut extra_use_line = String::new();
    // a linked-list problem
    if code.contains("pub struct ListNode") {
        extra_use_line.push_str("\nuse super::util::linked_list::{ListNode, to_list};")
    }
    if code.contains("pub struct TreeNode") {
        extra_use_line.push_str("\nuse super::util::tree::{TreeNode, to_tree};")
    }
    if code.contains("pub struct Point") {
        extra_use_line.push_str("\nuse super::util::point::Point;")
    }
    extra_use_line
}

fn build_desc(content: &str) -> String {
    // TODO: fix this shit
    content
        .replace("<strong>", "")
        .replace("</strong>", "")
        .replace("<em>", "")
        .replace("</em>", "")
        .replace("</p>", "")
        .replace("<p>", "")
        .replace("<b>", "")
        .replace("</b>", "")
        .replace("<pre>", "")
        .replace("</pre>", "")
        .replace("<ul>", "")
        .replace("</ul>", "")
        .replace("<li>", "")
        .replace("</li>", "")
        .replace("<code>", "")
        .replace("</code>", "")
        .replace("<i>", "")
        .replace("</i>", "")
        .replace("<sub>", "")
        .replace("</sub>", "")
        .replace("</sup>", "")
        .replace("<sup>", "^")
        .replace("&nbsp;", " ")
        .replace("&gt;", ">")
        .replace("&lt;", "<")
        .replace("&quot;", "\"")
        .replace("&minus;", "-")
        .replace("&#39;", "'")
        .replace("\n\n", "\n")
        .replace("\n", "\n * ")
}
fn generate_random_id(except_ids: &[u32]) -> u32 {
    use rand::Rng;
    use std::fs;
    let mut rng = rand::thread_rng();
    loop {
        let res: u32 = rng.gen_range(1, LEETCODE_NUMS);
        if !except_ids.contains(&res) {
            return res;
        }
    }
}
